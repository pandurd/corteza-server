package store

// This file is auto-generated.
//
// Template:    pkg/codegen/assets/store_base.gen.go.tpl
// Definitions: store/actionlog.yaml
//
// Changes to this file may cause incorrect behavior and will be lost if
// the code is regenerated.
//
// Definitions file that controls how this file is generated:
//  - store/actionlog.yaml

import (
	"context"
	"github.com/cortezaproject/corteza-server/pkg/actionlog"
)

type (
	Actionlogs interface {
		SearchActionlogs(ctx context.Context, f actionlog.Filter) (actionlog.ActionSet, actionlog.Filter, error)
		CreateActionlog(ctx context.Context, rr ...*actionlog.Action) error
		UpdateActionlog(ctx context.Context, rr ...*actionlog.Action) error
		PartialActionlogUpdate(ctx context.Context, onlyColumns []string, rr ...*actionlog.Action) error
		RemoveActionlog(ctx context.Context, rr ...*actionlog.Action) error
		RemoveActionlogByID(ctx context.Context, ID uint64) error

		TruncateActionlogs(ctx context.Context) error

		// Extra functions

	}
)

// SearchActionlogs returns all matching Actionlogs from store
func SearchActionlogs(ctx context.Context, s Actionlogs, f actionlog.Filter) (actionlog.ActionSet, actionlog.Filter, error) {
	return s.SearchActionlogs(ctx, f)
}

// CreateActionlog creates one or more Actionlogs in store
func CreateActionlog(ctx context.Context, s Actionlogs, rr ...*actionlog.Action) error {
	return s.CreateActionlog(ctx, rr...)
}

// UpdateActionlog updates one or more (existing) Actionlogs in store
func UpdateActionlog(ctx context.Context, s Actionlogs, rr ...*actionlog.Action) error {
	return s.UpdateActionlog(ctx, rr...)
}

// PartialActionlogUpdate updates one or more existing Actionlogs in store
func PartialActionlogUpdate(ctx context.Context, s Actionlogs, onlyColumns []string, rr ...*actionlog.Action) error {
	return s.PartialActionlogUpdate(ctx, onlyColumns, rr...)
}

// RemoveActionlog removes one or more Actionlogs from store
func RemoveActionlog(ctx context.Context, s Actionlogs, rr ...*actionlog.Action) error {
	return s.RemoveActionlog(ctx, rr...)
}

// RemoveActionlogByID removes Actionlog from store
func RemoveActionlogByID(ctx context.Context, s Actionlogs, ID uint64) error {
	return s.RemoveActionlogByID(ctx, ID)
}

// TruncateActionlogs removes all Actionlogs from store
func TruncateActionlogs(ctx context.Context, s Actionlogs) error {
	return s.TruncateActionlogs(ctx)
}
