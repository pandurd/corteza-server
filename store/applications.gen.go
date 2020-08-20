package store

// This file is auto-generated.
//
// Template:    pkg/codegen/assets/store_base.gen.go.tpl
// Definitions: store/applications.yaml
//
// Changes to this file may cause incorrect behavior and will be lost if
// the code is regenerated.
//
// Definitions file that controls how this file is generated:
//  - store/applications.yaml

import (
	"context"
	"github.com/cortezaproject/corteza-server/system/types"
)

type (
	Applications interface {
		SearchApplications(ctx context.Context, f types.ApplicationFilter) (types.ApplicationSet, types.ApplicationFilter, error)
		LookupApplicationByID(ctx context.Context, id uint64) (*types.Application, error)
		CreateApplication(ctx context.Context, rr ...*types.Application) error
		UpdateApplication(ctx context.Context, rr ...*types.Application) error
		PartialApplicationUpdate(ctx context.Context, onlyColumns []string, rr ...*types.Application) error
		RemoveApplication(ctx context.Context, rr ...*types.Application) error
		RemoveApplicationByID(ctx context.Context, ID uint64) error

		TruncateApplications(ctx context.Context) error

		// Extra functions
		ApplicationMetrics(ctx context.Context) (*types.ApplicationMetrics, error)
	}
)

// SearchApplications returns all matching Applications from store
func SearchApplications(ctx context.Context, s Applications, f types.ApplicationFilter) (types.ApplicationSet, types.ApplicationFilter, error) {
	return s.SearchApplications(ctx, f)
}

// LookupApplicationByID searches for application by ID
//
// It returns application even if deleted
func LookupApplicationByID(ctx context.Context, s Applications, id uint64) (*types.Application, error) {
	return s.LookupApplicationByID(ctx, id)
}

// CreateApplication creates one or more Applications in store
func CreateApplication(ctx context.Context, s Applications, rr ...*types.Application) error {
	return s.CreateApplication(ctx, rr...)
}

// UpdateApplication updates one or more (existing) Applications in store
func UpdateApplication(ctx context.Context, s Applications, rr ...*types.Application) error {
	return s.UpdateApplication(ctx, rr...)
}

// PartialApplicationUpdate updates one or more existing Applications in store
func PartialApplicationUpdate(ctx context.Context, s Applications, onlyColumns []string, rr ...*types.Application) error {
	return s.PartialApplicationUpdate(ctx, onlyColumns, rr...)
}

// RemoveApplication removes one or more Applications from store
func RemoveApplication(ctx context.Context, s Applications, rr ...*types.Application) error {
	return s.RemoveApplication(ctx, rr...)
}

// RemoveApplicationByID removes Application from store
func RemoveApplicationByID(ctx context.Context, s Applications, ID uint64) error {
	return s.RemoveApplicationByID(ctx, ID)
}

// TruncateApplications removes all Applications from store
func TruncateApplications(ctx context.Context, s Applications) error {
	return s.TruncateApplications(ctx)
}

func ApplicationMetrics(ctx context.Context, s Applications) (*types.ApplicationMetrics, error) {
	return s.ApplicationMetrics(ctx)
}
