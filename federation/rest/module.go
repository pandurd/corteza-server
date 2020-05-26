package rest

import (
	"context"
	"io/ioutil"

	composeService "github.com/cortezaproject/corteza-server/compose/service"
	composeTypes "github.com/cortezaproject/corteza-server/compose/types"
	"github.com/cortezaproject/corteza-server/federation/config"
	"github.com/cortezaproject/corteza-server/federation/rest/request"
	"github.com/cortezaproject/corteza-server/federation/service"
	"github.com/cortezaproject/corteza-server/federation/types"
)

type (
	moduleSetPayload struct {
		Filter composeTypes.ModuleFilter `json:"filter"`
		Set    []*modulePayload          `json:"set"`
	}

	modulePayload struct {
		*types.FederatedModule

		Fields []*moduleFieldPayload `json:"fields"`

		CanGrant        bool `json:"canGrant"`
		CanUpdateModule bool `json:"canUpdateModule"`
		CanDeleteModule bool `json:"canDeleteModule"`
		CanCreateRecord bool `json:"canCreateRecord"`
		CanReadRecord   bool `json:"canReadRecord"`
		CanUpdateRecord bool `json:"canUpdateRecord"`
		CanDeleteRecord bool `json:"canDeleteRecord"`

		CanManageAutomationTriggers bool `json:"canManageAutomationTriggers"`
	}

	moduleFieldPayload struct {
		*composeTypes.ModuleField

		CanReadRecordValue   bool `json:"canReadRecordValue"`
		CanUpdateRecordValue bool `json:"canUpdateRecordValue"`
	}

	Module struct {
		// module service.ModuleService
		ac  moduleAccessController
		cms composeService.ModuleService
		cns composeService.NamespaceService
	}

	moduleAccessController interface {
		CanGrant(context.Context) bool

		CanUpdateModule(context.Context, *composeTypes.Module) bool
		CanDeleteModule(context.Context, *composeTypes.Module) bool
		CanCreateRecord(context.Context, *composeTypes.Module) bool
		CanReadRecord(context.Context, *composeTypes.Module) bool
		CanUpdateRecord(context.Context, *composeTypes.Module) bool
		CanDeleteRecord(context.Context, *composeTypes.Module) bool

		CanReadRecordValue(context.Context, *composeTypes.ModuleField) bool
		CanUpdateRecordValue(context.Context, *composeTypes.ModuleField) bool

		CanManageAutomationTriggersOnModule(context.Context, *composeTypes.Module) bool
	}
)

func (Module) New() *Module {
	return &Module{
		// module: service.DefaultModule,
		ac:  service.DefaultAccessControl,
		cms: service.ComposeModuleService,
		cns: service.ComposeNamespaceService,
	}
}

func (ctrl *Module) List(ctx context.Context, r *request.ModuleList) (interface{}, error) {

	filter := composeTypes.ModuleFilter{}

	dat, err := ioutil.ReadFile("/home/wrk/Projects/corteza/corteza-server/federation/config/yaml/federation.yaml")

	if err != nil {
		return nil, err
	}

	parser := config.Parser{
		Config: dat,
	}
	federatedModuleList, err := parser.Modules()

	if err != nil {
		return nil, err
	}

	set := composeTypes.ModuleSet{}

	for _, module := range federatedModuleList.Modules {
		m, err := ctrl.cms.With(ctx).FindByHandle(module.Namespace.ID, module.Handle)
		if err != nil {
			panic(err)
		}

		set = append(set, m)
	}

	remappedList := []*types.FederatedModule{}
	namespaceList := []uint64{}
	namespaceInfo := map[uint64]*composeTypes.Namespace{}

	s, _, err := ctrl.cns.With(ctx).Find(composeTypes.NamespaceFilter{})

	for _, ns := range s {
		namespaceInfo[ns.ID] = ns
	}

	for _, m := range set {
		remappedModule := &types.FederatedModule{
			Module: composeTypes.Module{
				ID:          m.ID,
				Handle:      m.Handle,
				Name:        m.Name,
				Meta:        m.Meta,
				Fields:      m.Fields,
				NamespaceID: m.NamespaceID,
				CreatedAt:   m.CreatedAt,
			},
		}

		if namespaceInfo[m.NamespaceID] != nil {
			remappedModule.Namespace = *namespaceInfo[m.NamespaceID]
		}

		remappedList = append(remappedList, remappedModule)
		namespaceList = append(namespaceList, m.NamespaceID)
	}

	return ctrl.makeFilterPayload(ctx, remappedList, filter, err)
}

func (ctrl Module) makePayload(ctx context.Context, m *types.FederatedModule, err error) (*modulePayload, error) {
	if err != nil || m == nil {
		return nil, err
	}

	mfp, err := ctrl.makeFieldsPayload(ctx, m)
	if err != nil {
		return nil, err
	}

	return &modulePayload{
		FederatedModule: m,

		Fields: mfp,

		CanGrant: ctrl.ac.CanGrant(ctx),
	}, nil
}

func (ctrl Module) makeFieldsPayload(ctx context.Context, m *types.FederatedModule) (out []*moduleFieldPayload, err error) {
	out = make([]*moduleFieldPayload, len(m.Fields))

	for i, f := range m.Fields {
		out[i] = &moduleFieldPayload{
			ModuleField: f,

			CanReadRecordValue:   ctrl.ac.CanReadRecordValue(ctx, f),
			CanUpdateRecordValue: ctrl.ac.CanUpdateRecordValue(ctx, f),
		}
	}

	return
}

func (ctrl Module) makeFilterPayload(ctx context.Context, nn types.FederatedModuleSet, f composeTypes.ModuleFilter, err error) (*moduleSetPayload, error) {
	if err != nil {
		return nil, err
	}

	msp := &moduleSetPayload{Filter: f, Set: make([]*modulePayload, len(nn))}

	for i := range nn {
		msp.Set[i], _ = ctrl.makePayload(ctx, nn[i], nil)
	}

	return msp, nil
}
