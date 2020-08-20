package store

// This file is auto-generated.
//
// Template:    pkg/codegen/assets/store_base.gen.go.tpl
// Definitions: store/compose_modules.yaml
//
// Changes to this file may cause incorrect behavior and will be lost if
// the code is regenerated.
//
// Definitions file that controls how this file is generated:
//  - store/compose_modules.yaml

import (
	"context"
	"github.com/cortezaproject/corteza-server/compose/types"
)

type (
	ComposeModules interface {
		SearchComposeModules(ctx context.Context, f types.ModuleFilter) (types.ModuleSet, types.ModuleFilter, error)
		LookupComposeModuleByNamespaceIDHandle(ctx context.Context, namespace_id uint64, handle string) (*types.Module, error)
		LookupComposeModuleByNamespaceIDName(ctx context.Context, namespace_id uint64, name string) (*types.Module, error)
		LookupComposeModuleByID(ctx context.Context, id uint64) (*types.Module, error)
		CreateComposeModule(ctx context.Context, rr ...*types.Module) error
		UpdateComposeModule(ctx context.Context, rr ...*types.Module) error
		PartialComposeModuleUpdate(ctx context.Context, onlyColumns []string, rr ...*types.Module) error
		RemoveComposeModule(ctx context.Context, rr ...*types.Module) error
		RemoveComposeModuleByID(ctx context.Context, ID uint64) error

		TruncateComposeModules(ctx context.Context) error

		// Extra functions

	}
)

// SearchComposeModules returns all matching ComposeModules from store
func SearchComposeModules(ctx context.Context, s ComposeModules, f types.ModuleFilter) (types.ModuleSet, types.ModuleFilter, error) {
	return s.SearchComposeModules(ctx, f)
}

// LookupComposeModuleByNamespaceIDHandle searches for compose module by handle (case-insensitive)
func LookupComposeModuleByNamespaceIDHandle(ctx context.Context, s ComposeModules, namespace_id uint64, handle string) (*types.Module, error) {
	return s.LookupComposeModuleByNamespaceIDHandle(ctx, namespace_id, handle)
}

// LookupComposeModuleByNamespaceIDName searches for compose module by name (case-insensitive)
func LookupComposeModuleByNamespaceIDName(ctx context.Context, s ComposeModules, namespace_id uint64, name string) (*types.Module, error) {
	return s.LookupComposeModuleByNamespaceIDName(ctx, namespace_id, name)
}

// LookupComposeModuleByID searches for compose module by ID
//
// It returns compose module even if deleted
func LookupComposeModuleByID(ctx context.Context, s ComposeModules, id uint64) (*types.Module, error) {
	return s.LookupComposeModuleByID(ctx, id)
}

// CreateComposeModule creates one or more ComposeModules in store
func CreateComposeModule(ctx context.Context, s ComposeModules, rr ...*types.Module) error {
	return s.CreateComposeModule(ctx, rr...)
}

// UpdateComposeModule updates one or more (existing) ComposeModules in store
func UpdateComposeModule(ctx context.Context, s ComposeModules, rr ...*types.Module) error {
	return s.UpdateComposeModule(ctx, rr...)
}

// PartialComposeModuleUpdate updates one or more existing ComposeModules in store
func PartialComposeModuleUpdate(ctx context.Context, s ComposeModules, onlyColumns []string, rr ...*types.Module) error {
	return s.PartialComposeModuleUpdate(ctx, onlyColumns, rr...)
}

// RemoveComposeModule removes one or more ComposeModules from store
func RemoveComposeModule(ctx context.Context, s ComposeModules, rr ...*types.Module) error {
	return s.RemoveComposeModule(ctx, rr...)
}

// RemoveComposeModuleByID removes ComposeModule from store
func RemoveComposeModuleByID(ctx context.Context, s ComposeModules, ID uint64) error {
	return s.RemoveComposeModuleByID(ctx, ID)
}

// TruncateComposeModules removes all ComposeModules from store
func TruncateComposeModules(ctx context.Context, s ComposeModules) error {
	return s.TruncateComposeModules(ctx)
}
