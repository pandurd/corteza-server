package rest

import (
	"context"
	"github.com/cortezaproject/corteza-server/pkg/actionlog"
	"github.com/cortezaproject/corteza-server/pkg/payload"
	"github.com/cortezaproject/corteza-server/pkg/rh"
	"github.com/cortezaproject/corteza-server/store"
	"github.com/cortezaproject/corteza-server/system/rest/request"
	"github.com/cortezaproject/corteza-server/system/service"
	"github.com/cortezaproject/corteza-server/system/types"
)

type (
	Actionlog struct {
		actionSvc actionlog.Recorder
		userSvc   service.UserService
	}

	// Extend actionlog.Action so we can
	// provide user's email
	actionlogActionPayload struct {
		*actionlog.Action
		Actor string `json:"actor,omitempty"`
	}

	actionlogPayload struct {
		Filter actionlog.Filter          `json:"filter"`
		Set    []*actionlogActionPayload `json:"set"`
	}
)

func (Actionlog) New() *Actionlog {
	return &Actionlog{
		actionSvc: service.DefaultActionlog,
		userSvc:   service.DefaultUser,
	}
}

func (ctrl *Actionlog) List(ctx context.Context, r *request.ActionlogList) (interface{}, error) {
	var (
		err error
		f   = actionlog.Filter{
			From:     r.From,
			To:       r.To,
			ActorID:  payload.ParseUint64s(r.ActorID),
			Resource: r.Resource,
			Action:   r.Action,
		}
	)

	if f.Paging, err = store.NewPaging(r.Limit, r.PageCursor); err != nil {
		return nil, err
	}

	ee, f, err := ctrl.actionSvc.Find(ctx, f)

	return ctrl.makeFilterPayload(ctx, ee, f, err)
}

func (ctrl Actionlog) makeFilterPayload(ctx context.Context, ee []*actionlog.Action, f actionlog.Filter, err error) (*actionlogPayload, error) {
	if err != nil {
		return nil, err
	}

	var (
		pp = make([]*actionlogActionPayload, len(ee))
	)

	// Remap events to payload structs
	for e := range ee {
		pp[e] = &actionlogActionPayload{Action: ee[e]}
	}

	err = ctrl.userSvc.With(ctx).Preloader(
		func(c chan uint64) {
			for e := range ee {
				c <- ee[e].ActorID
			}

			close(c)
		},
		types.UserFilter{
			Deleted:   rh.FilterStateInclusive,
			Suspended: rh.FilterStateInclusive,
		},
		func(u *types.User) error {
			for p := range pp {
				if pp[p].ActorID == u.ID {
					pp[p].Actor = u.Name
					if pp[p].Actor == "" {
						pp[p].Actor = u.Email
					}
				}
			}

			return nil
		},
	)

	if err != nil {
		return nil, err
	}

	return &actionlogPayload{Filter: f, Set: pp}, nil
}
