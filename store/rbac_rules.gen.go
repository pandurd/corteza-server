package store

// This file is auto-generated.
//
// Template:    pkg/codegen/assets/store_base.gen.go.tpl
// Definitions: store/rbac_rules.yaml
//
// Changes to this file may cause incorrect behavior and will be lost if
// the code is regenerated.
//
// Definitions file that controls how this file is generated:
//  - store/rbac_rules.yaml

import (
	"context"
	"github.com/cortezaproject/corteza-server/pkg/permissions"
)

type (
	RbacRules interface {
		SearchRbacRules(ctx context.Context, f permissions.RuleFilter) (permissions.RuleSet, permissions.RuleFilter, error)
		CreateRbacRule(ctx context.Context, rr ...*permissions.Rule) error
		UpdateRbacRule(ctx context.Context, rr ...*permissions.Rule) error
		PartialRbacRuleUpdate(ctx context.Context, onlyColumns []string, rr ...*permissions.Rule) error
		RemoveRbacRule(ctx context.Context, rr ...*permissions.Rule) error
		RemoveRbacRuleByRoleIDResourceOperation(ctx context.Context, roleID uint64, resource string, operation string) error

		TruncateRbacRules(ctx context.Context) error

		// Extra functions

	}
)

// SearchRbacRules returns all matching RbacRules from store
func SearchRbacRules(ctx context.Context, s RbacRules, f permissions.RuleFilter) (permissions.RuleSet, permissions.RuleFilter, error) {
	return s.SearchRbacRules(ctx, f)
}

// CreateRbacRule creates one or more RbacRules in store
func CreateRbacRule(ctx context.Context, s RbacRules, rr ...*permissions.Rule) error {
	return s.CreateRbacRule(ctx, rr...)
}

// UpdateRbacRule updates one or more (existing) RbacRules in store
func UpdateRbacRule(ctx context.Context, s RbacRules, rr ...*permissions.Rule) error {
	return s.UpdateRbacRule(ctx, rr...)
}

// PartialRbacRuleUpdate updates one or more existing RbacRules in store
func PartialRbacRuleUpdate(ctx context.Context, s RbacRules, onlyColumns []string, rr ...*permissions.Rule) error {
	return s.PartialRbacRuleUpdate(ctx, onlyColumns, rr...)
}

// RemoveRbacRule removes one or more RbacRules from store
func RemoveRbacRule(ctx context.Context, s RbacRules, rr ...*permissions.Rule) error {
	return s.RemoveRbacRule(ctx, rr...)
}

// RemoveRbacRuleByRoleIDResourceOperation removes RbacRule from store
func RemoveRbacRuleByRoleIDResourceOperation(ctx context.Context, s RbacRules, roleID uint64, resource string, operation string) error {
	return s.RemoveRbacRuleByRoleIDResourceOperation(ctx, roleID, resource, operation)
}

// TruncateRbacRules removes all RbacRules from store
func TruncateRbacRules(ctx context.Context, s RbacRules) error {
	return s.TruncateRbacRules(ctx)
}
