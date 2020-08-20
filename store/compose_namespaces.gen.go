package store

// This file is auto-generated.
//
// Template:    pkg/codegen/assets/store_base.gen.go.tpl
// Definitions: store/compose_namespaces.yaml
//
// Changes to this file may cause incorrect behavior and will be lost if
// the code is regenerated.
//
// Definitions file that controls how this file is generated:
//  - store/compose_namespaces.yaml

import (
	"context"
	"github.com/cortezaproject/corteza-server/compose/types"
)

type (
	ComposeNamespaces interface {
		SearchComposeNamespaces(ctx context.Context, f types.NamespaceFilter) (types.NamespaceSet, types.NamespaceFilter, error)
		LookupComposeNamespaceBySlug(ctx context.Context, slug string) (*types.Namespace, error)
		LookupComposeNamespaceByID(ctx context.Context, id uint64) (*types.Namespace, error)
		CreateComposeNamespace(ctx context.Context, rr ...*types.Namespace) error
		UpdateComposeNamespace(ctx context.Context, rr ...*types.Namespace) error
		PartialComposeNamespaceUpdate(ctx context.Context, onlyColumns []string, rr ...*types.Namespace) error
		RemoveComposeNamespace(ctx context.Context, rr ...*types.Namespace) error
		RemoveComposeNamespaceByID(ctx context.Context, ID uint64) error

		TruncateComposeNamespaces(ctx context.Context) error

		// Extra functions

	}
)

// SearchComposeNamespaces returns all matching ComposeNamespaces from store
func SearchComposeNamespaces(ctx context.Context, s ComposeNamespaces, f types.NamespaceFilter) (types.NamespaceSet, types.NamespaceFilter, error) {
	return s.SearchComposeNamespaces(ctx, f)
}

// LookupComposeNamespaceBySlug searches for namespace by slug (case-insensitive)
func LookupComposeNamespaceBySlug(ctx context.Context, s ComposeNamespaces, slug string) (*types.Namespace, error) {
	return s.LookupComposeNamespaceBySlug(ctx, slug)
}

// LookupComposeNamespaceByID searches for compose namespace by ID
//
// It returns compose namespace even if deleted
func LookupComposeNamespaceByID(ctx context.Context, s ComposeNamespaces, id uint64) (*types.Namespace, error) {
	return s.LookupComposeNamespaceByID(ctx, id)
}

// CreateComposeNamespace creates one or more ComposeNamespaces in store
func CreateComposeNamespace(ctx context.Context, s ComposeNamespaces, rr ...*types.Namespace) error {
	return s.CreateComposeNamespace(ctx, rr...)
}

// UpdateComposeNamespace updates one or more (existing) ComposeNamespaces in store
func UpdateComposeNamespace(ctx context.Context, s ComposeNamespaces, rr ...*types.Namespace) error {
	return s.UpdateComposeNamespace(ctx, rr...)
}

// PartialComposeNamespaceUpdate updates one or more existing ComposeNamespaces in store
func PartialComposeNamespaceUpdate(ctx context.Context, s ComposeNamespaces, onlyColumns []string, rr ...*types.Namespace) error {
	return s.PartialComposeNamespaceUpdate(ctx, onlyColumns, rr...)
}

// RemoveComposeNamespace removes one or more ComposeNamespaces from store
func RemoveComposeNamespace(ctx context.Context, s ComposeNamespaces, rr ...*types.Namespace) error {
	return s.RemoveComposeNamespace(ctx, rr...)
}

// RemoveComposeNamespaceByID removes ComposeNamespace from store
func RemoveComposeNamespaceByID(ctx context.Context, s ComposeNamespaces, ID uint64) error {
	return s.RemoveComposeNamespaceByID(ctx, ID)
}

// TruncateComposeNamespaces removes all ComposeNamespaces from store
func TruncateComposeNamespaces(ctx context.Context, s ComposeNamespaces) error {
	return s.TruncateComposeNamespaces(ctx)
}
