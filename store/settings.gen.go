package store

// This file is auto-generated.
//
// Template:    pkg/codegen/assets/store_base.gen.go.tpl
// Definitions: store/settings.yaml
//
// Changes to this file may cause incorrect behavior and will be lost if
// the code is regenerated.
//
// Definitions file that controls how this file is generated:
//  - store/settings.yaml

import (
	"context"
	"github.com/cortezaproject/corteza-server/system/types"
)

type (
	Settings interface {
		SearchSettings(ctx context.Context, f types.SettingsFilter) (types.SettingValueSet, types.SettingsFilter, error)
		LookupSettingByNameOwnedBy(ctx context.Context, name string, owned_by uint64) (*types.SettingValue, error)
		CreateSetting(ctx context.Context, rr ...*types.SettingValue) error
		UpdateSetting(ctx context.Context, rr ...*types.SettingValue) error
		PartialSettingUpdate(ctx context.Context, onlyColumns []string, rr ...*types.SettingValue) error
		RemoveSetting(ctx context.Context, rr ...*types.SettingValue) error
		RemoveSettingByNameOwnedBy(ctx context.Context, name string, ownedBy uint64) error

		TruncateSettings(ctx context.Context) error

		// Extra functions

	}
)

// SearchSettings returns all matching Settings from store
func SearchSettings(ctx context.Context, s Settings, f types.SettingsFilter) (types.SettingValueSet, types.SettingsFilter, error) {
	return s.SearchSettings(ctx, f)
}

// LookupSettingByNameOwnedBy searches for settings by name and owner
func LookupSettingByNameOwnedBy(ctx context.Context, s Settings, name string, owned_by uint64) (*types.SettingValue, error) {
	return s.LookupSettingByNameOwnedBy(ctx, name, owned_by)
}

// CreateSetting creates one or more Settings in store
func CreateSetting(ctx context.Context, s Settings, rr ...*types.SettingValue) error {
	return s.CreateSetting(ctx, rr...)
}

// UpdateSetting updates one or more (existing) Settings in store
func UpdateSetting(ctx context.Context, s Settings, rr ...*types.SettingValue) error {
	return s.UpdateSetting(ctx, rr...)
}

// PartialSettingUpdate updates one or more existing Settings in store
func PartialSettingUpdate(ctx context.Context, s Settings, onlyColumns []string, rr ...*types.SettingValue) error {
	return s.PartialSettingUpdate(ctx, onlyColumns, rr...)
}

// RemoveSetting removes one or more Settings from store
func RemoveSetting(ctx context.Context, s Settings, rr ...*types.SettingValue) error {
	return s.RemoveSetting(ctx, rr...)
}

// RemoveSettingByNameOwnedBy removes Setting from store
func RemoveSettingByNameOwnedBy(ctx context.Context, s Settings, name string, ownedBy uint64) error {
	return s.RemoveSettingByNameOwnedBy(ctx, name, ownedBy)
}

// TruncateSettings removes all Settings from store
func TruncateSettings(ctx context.Context, s Settings) error {
	return s.TruncateSettings(ctx)
}
