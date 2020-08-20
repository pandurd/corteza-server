package store

// This file is auto-generated.
//
// Template:    pkg/codegen/assets/store_base.gen.go.tpl
// Definitions: store/credentials.yaml
//
// Changes to this file may cause incorrect behavior and will be lost if
// the code is regenerated.
//
// Definitions file that controls how this file is generated:
//  - store/credentials.yaml

import (
	"context"
	"github.com/cortezaproject/corteza-server/system/types"
)

type (
	Credentials interface {
		SearchCredentials(ctx context.Context, f types.CredentialsFilter) (types.CredentialsSet, types.CredentialsFilter, error)
		LookupCredentialsByID(ctx context.Context, id uint64) (*types.Credentials, error)
		CreateCredentials(ctx context.Context, rr ...*types.Credentials) error
		UpdateCredentials(ctx context.Context, rr ...*types.Credentials) error
		PartialCredentialsUpdate(ctx context.Context, onlyColumns []string, rr ...*types.Credentials) error
		RemoveCredentials(ctx context.Context, rr ...*types.Credentials) error
		RemoveCredentialsByID(ctx context.Context, ID uint64) error

		TruncateCredentials(ctx context.Context) error

		// Extra functions

	}
)

// SearchCredentials returns all matching Credentials from store
func SearchCredentials(ctx context.Context, s Credentials, f types.CredentialsFilter) (types.CredentialsSet, types.CredentialsFilter, error) {
	return s.SearchCredentials(ctx, f)
}

// LookupCredentialsByID searches for credentials by ID
//
// It returns credentials even if deleted
func LookupCredentialsByID(ctx context.Context, s Credentials, id uint64) (*types.Credentials, error) {
	return s.LookupCredentialsByID(ctx, id)
}

// CreateCredentials creates one or more Credentials in store
func CreateCredentials(ctx context.Context, s Credentials, rr ...*types.Credentials) error {
	return s.CreateCredentials(ctx, rr...)
}

// UpdateCredentials updates one or more (existing) Credentials in store
func UpdateCredentials(ctx context.Context, s Credentials, rr ...*types.Credentials) error {
	return s.UpdateCredentials(ctx, rr...)
}

// PartialCredentialsUpdate updates one or more existing Credentials in store
func PartialCredentialsUpdate(ctx context.Context, s Credentials, onlyColumns []string, rr ...*types.Credentials) error {
	return s.PartialCredentialsUpdate(ctx, onlyColumns, rr...)
}

// RemoveCredentials removes one or more Credentials from store
func RemoveCredentials(ctx context.Context, s Credentials, rr ...*types.Credentials) error {
	return s.RemoveCredentials(ctx, rr...)
}

// RemoveCredentialsByID removes Credentials from store
func RemoveCredentialsByID(ctx context.Context, s Credentials, ID uint64) error {
	return s.RemoveCredentialsByID(ctx, ID)
}

// TruncateCredentials removes all Credentials from store
func TruncateCredentials(ctx context.Context, s Credentials) error {
	return s.TruncateCredentials(ctx)
}
