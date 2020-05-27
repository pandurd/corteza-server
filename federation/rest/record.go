package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/davecgh/go-spew/spew"
	"github.com/pkg/errors"

	composeService "github.com/cortezaproject/corteza-server/compose/service"
	composeTypes "github.com/cortezaproject/corteza-server/compose/types"
	"github.com/cortezaproject/corteza-server/federation/rest/request"
	"github.com/cortezaproject/corteza-server/federation/service"
	"github.com/cortezaproject/corteza-server/federation/types"
	"github.com/cortezaproject/corteza-server/pkg/rh"
)

var _ = errors.Wrap

type (
	recordPayload struct {
		*types.FederatedRecord
	}

	recordSetPayload struct {
		Filter composeTypes.RecordFilter `json:"filter"`
		Set    []*recordPayload          `json:"set"`
	}

	Record struct {
		crs composeService.RecordService
		cms composeService.ModuleService
		cns composeService.NamespaceService
	}

	// recordAccessController interface {
	// 	CanUpdateRecord(context.Context, *types.Module) bool
	// 	CanDeleteRecord(context.Context, *types.Module) bool
	// }
)

func (Record) New() *Record {
	return &Record{
		crs: service.ComposeRecordService,
		cms: service.ComposeModuleService,
		cns: service.ComposeNamespaceService,
	}
}

func (ctrl *Record) List(ctx context.Context, r *request.RecordList) (interface{}, error) {
	var (
		// m   *composeTypes.Module
		err error

		rf = composeTypes.RecordFilter{
			NamespaceID: r.NamespaceID,
			ModuleID:    r.ModuleID,
			Sort:        r.Sort,

			Deleted: rh.FilterState(r.Deleted),

			PageFilter: rh.Paging(r),
		}
	)

	// if _, err = ctrl.cms.With(ctx).FindByID(r.NamespaceID, r.ModuleID); err != nil {
	// 	return nil, err
	// }

	if r.Query != "" {
		// Query param takes preference
		rf.Query = r.Query
	} else if r.Filter != "" {
		// Backward compatibility
		// Filter param is deprecated
		rf.Query = r.Filter
	}

	rf.Deleted = 1

	if !r.LastSynced.IsZero() {
		mFormat := r.LastSynced.Format("2006-01-02 15:04:05")
		rf.Query = fmt.Sprintf("createdAt >= '%s' OR updatedAt >= '%s' OR deletedAt >= '%s'", mFormat, mFormat, mFormat)
	}

	spew.Dump("QQQQQ", rf.Query)

	rr, filter, err := ctrl.crs.With(ctx).Find(rf)

	remappedList := []*types.FederatedRecord{}

	for _, r := range rr {
		remappedRecord := &types.FederatedRecord{
			ID:          r.ID,
			ModuleID:    r.ModuleID,
			NamespaceID: r.NamespaceID,
			CreatedAt:   r.CreatedAt,
			DeletedAt:   r.DeletedAt,
			UpdatedAt:   r.UpdatedAt,
			CreatedBy:   r.CreatedBy,
			DeletedBy:   r.DeletedBy,
			UpdatedBy:   r.UpdatedBy,
			Values:      r.Values,
		}

		remappedList = append(remappedList, remappedRecord)
	}

	return ctrl.makeFilterPayload(ctx, remappedList, filter, err)
}

func (ctrl Record) makePayload(ctx context.Context, r *types.FederatedRecord, err error) (*recordPayload, error) {
	if err != nil || r == nil {
		return nil, err
	}

	return &recordPayload{
		FederatedRecord: r,
	}, nil
}

func (ctrl Record) makeFilterPayload(ctx context.Context, rr types.FederatedRecordSet, f composeTypes.RecordFilter, err error) (*recordSetPayload, error) {
	if err != nil {
		return nil, err
	}

	modp := &recordSetPayload{Filter: f, Set: make([]*recordPayload, len(rr))}

	for i := range rr {
		modp.Set[i], _ = ctrl.makePayload(ctx, rr[i], nil)
	}

	return modp, nil
}

// Special care for record validation errors
//
// We need to return a bit different format of response
// with all details that were collected through validation
func (ctrl Record) handleValidationError(rve *composeTypes.RecordValueErrorSet) interface{} {
	return func(w http.ResponseWriter, _ *http.Request) {
		rval := struct {
			Error struct {
				Message string                          `json:"message"`
				Details []composeTypes.RecordValueError `json:"details,omitempty"`
			} `json:"error"`
		}{}

		rval.Error.Message = rve.Error()
		rval.Error.Details = rve.Set

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(rval)
	}
}

func (ctrl Record) DecodeFilterPayload(ctx context.Context, payload []byte) (types.FederatedRecordSet, error) {
	// federatedRecordPayload := &recordSetPayload{}

	type Aux struct {
		// Response map[string]interface{} `json:"response"`
		Response struct {
			Filter interface{}              `json:"filter"`
			Set    types.FederatedRecordSet `json:"set"`
		} `json:"response"`
	}

	aux := Aux{}

	err := json.Unmarshal(payload, &aux)

	return aux.Response.Set, err
}
