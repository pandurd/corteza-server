package store

// This file is auto-generated.
//
// Template:    pkg/codegen/assets/store_base.gen.go.tpl
// Definitions: store/compose_module_fields.yaml
//
// Changes to this file may cause incorrect behavior and will be lost if
// the code is regenerated.
//
// Definitions file that controls how this file is generated:
//  - store/compose_module_fields.yaml

import (
	"context"
	"github.com/cortezaproject/corteza-server/compose/types"
)

type (
	ComposeModuleFields interface {
		SearchComposeModuleFields(ctx context.Context, f types.ModuleFieldFilter) (types.ModuleFieldSet, types.ModuleFieldFilter, error)
		CreateComposeModuleField(ctx context.Context, rr ...*types.ModuleField) error
		UpdateComposeModuleField(ctx context.Context, rr ...*types.ModuleField) error
		PartialComposeModuleFieldUpdate(ctx context.Context, onlyColumns []string, rr ...*types.ModuleField) error
		RemoveComposeModuleField(ctx context.Context, rr ...*types.ModuleField) error
		RemoveComposeModuleFieldByID(ctx context.Context, ID uint64) error

		TruncateComposeModuleFields(ctx context.Context) error

		// Extra functions

	}
)

// SearchComposeModuleFields returns all matching ComposeModuleFields from store
func SearchComposeModuleFields(ctx context.Context, s ComposeModuleFields, f types.ModuleFieldFilter) (types.ModuleFieldSet, types.ModuleFieldFilter, error) {
	return s.SearchComposeModuleFields(ctx, f)
}

// CreateComposeModuleField creates one or more ComposeModuleFields in store
func CreateComposeModuleField(ctx context.Context, s ComposeModuleFields, rr ...*types.ModuleField) error {
	return s.CreateComposeModuleField(ctx, rr...)
}

// UpdateComposeModuleField updates one or more (existing) ComposeModuleFields in store
func UpdateComposeModuleField(ctx context.Context, s ComposeModuleFields, rr ...*types.ModuleField) error {
	return s.UpdateComposeModuleField(ctx, rr...)
}

// PartialComposeModuleFieldUpdate updates one or more existing ComposeModuleFields in store
func PartialComposeModuleFieldUpdate(ctx context.Context, s ComposeModuleFields, onlyColumns []string, rr ...*types.ModuleField) error {
	return s.PartialComposeModuleFieldUpdate(ctx, onlyColumns, rr...)
}

// RemoveComposeModuleField removes one or more ComposeModuleFields from store
func RemoveComposeModuleField(ctx context.Context, s ComposeModuleFields, rr ...*types.ModuleField) error {
	return s.RemoveComposeModuleField(ctx, rr...)
}

// RemoveComposeModuleFieldByID removes ComposeModuleField from store
func RemoveComposeModuleFieldByID(ctx context.Context, s ComposeModuleFields, ID uint64) error {
	return s.RemoveComposeModuleFieldByID(ctx, ID)
}

// TruncateComposeModuleFields removes all ComposeModuleFields from store
func TruncateComposeModuleFields(ctx context.Context, s ComposeModuleFields) error {
	return s.TruncateComposeModuleFields(ctx)
}
