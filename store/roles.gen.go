package store

// This file is auto-generated.
//
// Template:    pkg/codegen/assets/store_base.gen.go.tpl
// Definitions: store/roles.yaml
//
// Changes to this file may cause incorrect behavior and will be lost if
// the code is regenerated.
//
// Definitions file that controls how this file is generated:
//  - store/roles.yaml

import (
	"context"
	"github.com/cortezaproject/corteza-server/system/types"
)

type (
	Roles interface {
		SearchRoles(ctx context.Context, f types.RoleFilter) (types.RoleSet, types.RoleFilter, error)
		LookupRoleByID(ctx context.Context, id uint64) (*types.Role, error)
		LookupRoleByHandle(ctx context.Context, handle string) (*types.Role, error)
		LookupRoleByName(ctx context.Context, name string) (*types.Role, error)
		CreateRole(ctx context.Context, rr ...*types.Role) error
		UpdateRole(ctx context.Context, rr ...*types.Role) error
		PartialRoleUpdate(ctx context.Context, onlyColumns []string, rr ...*types.Role) error
		RemoveRole(ctx context.Context, rr ...*types.Role) error
		RemoveRoleByID(ctx context.Context, ID uint64) error

		TruncateRoles(ctx context.Context) error

		// Extra functions
		RoleMetrics(ctx context.Context) (*types.RoleMetrics, error)
	}
)

// SearchRoles returns all matching Roles from store
func SearchRoles(ctx context.Context, s Roles, f types.RoleFilter) (types.RoleSet, types.RoleFilter, error) {
	return s.SearchRoles(ctx, f)
}

// LookupRoleByID searches for role by ID
//
// It returns role even if deleted or suspended
func LookupRoleByID(ctx context.Context, s Roles, id uint64) (*types.Role, error) {
	return s.LookupRoleByID(ctx, id)
}

// LookupRoleByHandle searches for role by its handle
//
// It returns only valid roles (not deleted, not archived)
func LookupRoleByHandle(ctx context.Context, s Roles, handle string) (*types.Role, error) {
	return s.LookupRoleByHandle(ctx, handle)
}

// LookupRoleByName searches for role by its name
//
// It returns only valid roles (not deleted, not archived)
func LookupRoleByName(ctx context.Context, s Roles, name string) (*types.Role, error) {
	return s.LookupRoleByName(ctx, name)
}

// CreateRole creates one or more Roles in store
func CreateRole(ctx context.Context, s Roles, rr ...*types.Role) error {
	return s.CreateRole(ctx, rr...)
}

// UpdateRole updates one or more (existing) Roles in store
func UpdateRole(ctx context.Context, s Roles, rr ...*types.Role) error {
	return s.UpdateRole(ctx, rr...)
}

// PartialRoleUpdate updates one or more existing Roles in store
func PartialRoleUpdate(ctx context.Context, s Roles, onlyColumns []string, rr ...*types.Role) error {
	return s.PartialRoleUpdate(ctx, onlyColumns, rr...)
}

// RemoveRole removes one or more Roles from store
func RemoveRole(ctx context.Context, s Roles, rr ...*types.Role) error {
	return s.RemoveRole(ctx, rr...)
}

// RemoveRoleByID removes Role from store
func RemoveRoleByID(ctx context.Context, s Roles, ID uint64) error {
	return s.RemoveRoleByID(ctx, ID)
}

// TruncateRoles removes all Roles from store
func TruncateRoles(ctx context.Context, s Roles) error {
	return s.TruncateRoles(ctx)
}

func RoleMetrics(ctx context.Context, s Roles) (*types.RoleMetrics, error) {
	return s.RoleMetrics(ctx)
}
