package store

// This file is auto-generated.
//
// Template:    pkg/codegen/assets/store_base.gen.go.tpl
// Definitions: store/role_members.yaml
//
// Changes to this file may cause incorrect behavior and will be lost if
// the code is regenerated.
//
// Definitions file that controls how this file is generated:
//  - store/role_members.yaml

import (
	"context"
	"github.com/cortezaproject/corteza-server/system/types"
)

type (
	RoleMembers interface {
		SearchRoleMembers(ctx context.Context, f types.RoleMemberFilter) (types.RoleMemberSet, types.RoleMemberFilter, error)
		CreateRoleMember(ctx context.Context, rr ...*types.RoleMember) error
		UpdateRoleMember(ctx context.Context, rr ...*types.RoleMember) error
		PartialRoleMemberUpdate(ctx context.Context, onlyColumns []string, rr ...*types.RoleMember) error
		RemoveRoleMember(ctx context.Context, rr ...*types.RoleMember) error
		RemoveRoleMemberByUserIDRoleID(ctx context.Context, userID uint64, roleID uint64) error

		TruncateRoleMembers(ctx context.Context) error

		// Extra functions

	}
)

// SearchRoleMembers returns all matching RoleMembers from store
func SearchRoleMembers(ctx context.Context, s RoleMembers, f types.RoleMemberFilter) (types.RoleMemberSet, types.RoleMemberFilter, error) {
	return s.SearchRoleMembers(ctx, f)
}

// CreateRoleMember creates one or more RoleMembers in store
func CreateRoleMember(ctx context.Context, s RoleMembers, rr ...*types.RoleMember) error {
	return s.CreateRoleMember(ctx, rr...)
}

// UpdateRoleMember updates one or more (existing) RoleMembers in store
func UpdateRoleMember(ctx context.Context, s RoleMembers, rr ...*types.RoleMember) error {
	return s.UpdateRoleMember(ctx, rr...)
}

// PartialRoleMemberUpdate updates one or more existing RoleMembers in store
func PartialRoleMemberUpdate(ctx context.Context, s RoleMembers, onlyColumns []string, rr ...*types.RoleMember) error {
	return s.PartialRoleMemberUpdate(ctx, onlyColumns, rr...)
}

// RemoveRoleMember removes one or more RoleMembers from store
func RemoveRoleMember(ctx context.Context, s RoleMembers, rr ...*types.RoleMember) error {
	return s.RemoveRoleMember(ctx, rr...)
}

// RemoveRoleMemberByUserIDRoleID removes RoleMember from store
func RemoveRoleMemberByUserIDRoleID(ctx context.Context, s RoleMembers, userID uint64, roleID uint64) error {
	return s.RemoveRoleMemberByUserIDRoleID(ctx, userID, roleID)
}

// TruncateRoleMembers removes all RoleMembers from store
func TruncateRoleMembers(ctx context.Context, s RoleMembers) error {
	return s.TruncateRoleMembers(ctx)
}
