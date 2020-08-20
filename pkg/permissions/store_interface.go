package permissions

import (
	"context"
)

type (
	// All
	rbacRulesStore interface {
		SearchRbacRules(ctx context.Context, f RuleFilter) (RuleSet, RuleFilter, error)
		CreateRbacRule(ctx context.Context, rr ...*Rule) error
		UpdateRbacRule(ctx context.Context, rr ...*Rule) error
		RemoveRbacRule(ctx context.Context, rr ...*Rule) error
		TruncateRbacRules(ctx context.Context) error
	}
)
