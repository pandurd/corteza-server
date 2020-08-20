package rdbms

// This file is an auto-generated file
//
// Template:    pkg/codegen/assets/store_rdbms.gen.go.tpl
// Definitions: store/compose_module_fields.yaml
//
// Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated.

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/cortezaproject/corteza-server/compose/types"
	"github.com/cortezaproject/corteza-server/store"
)

// SearchComposeModuleFields returns all matching rows
//
// This function calls convertComposeModuleFieldFilter with the given
// types.ModuleFieldFilter and expects to receive a working squirrel.SelectBuilder
func (s Store) SearchComposeModuleFields(ctx context.Context, f types.ModuleFieldFilter) (types.ModuleFieldSet, types.ModuleFieldFilter, error) {
	var scap uint
	q, err := s.convertComposeModuleFieldFilter(f)
	if err != nil {
		return nil, f, err
	}

	if scap == 0 {
		scap = DefaultSliceCapacity
	}

	var (
		set = make([]*types.ModuleField, 0, scap)
		// Paging is disabled in definition yaml file
		// {search: {disablePaging:true}} and this allows
		// a much simpler row fetching logic
		fetch = func() error {
			var (
				res       *types.ModuleField
				rows, err = s.Query(ctx, q)
			)

			if err != nil {
				return err
			}

			for rows.Next() {
				if res, err = s.internalComposeModuleFieldRowScanner(rows, rows.Err()); err != nil {
					if cerr := rows.Close(); cerr != nil {
						return fmt.Errorf("could not close rows (%v) after scan error: %w", cerr, err)
					}

					return err
				}

				set = append(set, res)
			}

			return rows.Close()
		}
	)

	return set, f, s.config.ErrorHandler(fetch())
}

// CreateComposeModuleField creates one or more rows in compose_module_field table
func (s Store) CreateComposeModuleField(ctx context.Context, rr ...*types.ModuleField) (err error) {
	for _, res := range rr {
		err = s.Exec(ctx, s.InsertBuilder(s.ComposeModuleFieldTable()).SetMap(s.internalComposeModuleFieldEncoder(res)))
		if err != nil {
			return s.config.ErrorHandler(err)
		}
	}

	return
}

// UpdateComposeModuleField updates one or more existing rows in compose_module_field
func (s Store) UpdateComposeModuleField(ctx context.Context, rr ...*types.ModuleField) error {
	return s.config.ErrorHandler(s.PartialComposeModuleFieldUpdate(ctx, nil, rr...))
}

// PartialComposeModuleFieldUpdate updates one or more existing rows in compose_module_field
func (s Store) PartialComposeModuleFieldUpdate(ctx context.Context, onlyColumns []string, rr ...*types.ModuleField) (err error) {
	for _, res := range rr {
		err = s.ExecUpdateComposeModuleFields(
			ctx,
			squirrel.Eq{s.preprocessColumn("cmf.id", ""): s.preprocessValue(res.ID, "")},
			s.internalComposeModuleFieldEncoder(res).Skip("id").Only(onlyColumns...))
		if err != nil {
			return s.config.ErrorHandler(err)
		}
	}

	return
}

// RemoveComposeModuleField removes one or more rows from compose_module_field table
func (s Store) RemoveComposeModuleField(ctx context.Context, rr ...*types.ModuleField) (err error) {
	for _, res := range rr {
		err = s.Exec(ctx, s.DeleteBuilder(s.ComposeModuleFieldTable("cmf")).Where(squirrel.Eq{s.preprocessColumn("cmf.id", ""): s.preprocessValue(res.ID, "")}))
		if err != nil {
			return s.config.ErrorHandler(err)
		}
	}

	return nil
}

// RemoveComposeModuleFieldByID removes row from the compose_module_field table
func (s Store) RemoveComposeModuleFieldByID(ctx context.Context, ID uint64) error {
	return s.config.ErrorHandler(s.Exec(ctx, s.DeleteBuilder(s.ComposeModuleFieldTable("cmf")).Where(squirrel.Eq{s.preprocessColumn("cmf.id", ""): s.preprocessValue(ID, "")})))
}

// TruncateComposeModuleFields removes all rows from the compose_module_field table
func (s Store) TruncateComposeModuleFields(ctx context.Context) error {
	return s.config.ErrorHandler(s.Truncate(ctx, s.ComposeModuleFieldTable()))
}

// ExecUpdateComposeModuleFields updates all matched (by cnd) rows in compose_module_field with given data
func (s Store) ExecUpdateComposeModuleFields(ctx context.Context, cnd squirrel.Sqlizer, set store.Payload) error {
	return s.config.ErrorHandler(s.Exec(ctx, s.UpdateBuilder(s.ComposeModuleFieldTable("cmf")).Where(cnd).SetMap(set)))
}

// ComposeModuleFieldLookup prepares ComposeModuleField query and executes it,
// returning types.ModuleField (or error)
func (s Store) ComposeModuleFieldLookup(ctx context.Context, cnd squirrel.Sqlizer) (*types.ModuleField, error) {
	return s.internalComposeModuleFieldRowScanner(s.QueryRow(ctx, s.QueryComposeModuleFields().Where(cnd)))
}

func (s Store) internalComposeModuleFieldRowScanner(row rowScanner, err error) (*types.ModuleField, error) {
	if err != nil {
		return nil, err
	}

	var res = &types.ModuleField{}
	if _, has := s.config.RowScanners["composeModuleField"]; has {
		scanner := s.config.RowScanners["composeModuleField"].(func(rowScanner, *types.ModuleField) error)
		err = scanner(row, res)
	} else {
		err = row.Scan(
			&res.ID,
			&res.Name,
			&res.ModuleID,
			&res.Place,
			&res.Kind,
			&res.Label,
			&res.Options,
			&res.Private,
			&res.Required,
			&res.Visible,
			&res.Multi,
			&res.DefaultValue,
			&res.CreatedAt,
			&res.UpdatedAt,
			&res.DeletedAt,
		)
	}

	if err == sql.ErrNoRows {
		return nil, store.ErrNotFound
	}

	if err != nil {
		return nil, fmt.Errorf("could not scan db row for ComposeModuleField: %w", err)
	} else {
		return res, nil
	}
}

// QueryComposeModuleFields returns squirrel.SelectBuilder with set table and all columns
func (s Store) QueryComposeModuleFields() squirrel.SelectBuilder {
	return s.SelectBuilder(s.ComposeModuleFieldTable("cmf"), s.ComposeModuleFieldColumns("cmf")...)
}

// ComposeModuleFieldTable name of the db table
func (Store) ComposeModuleFieldTable(aa ...string) string {
	var alias string
	if len(aa) > 0 {
		alias = " AS " + aa[0]
	}

	return "compose_module_field" + alias
}

// ComposeModuleFieldColumns returns all defined table columns
//
// With optional string arg, all columns are returned aliased
func (Store) ComposeModuleFieldColumns(aa ...string) []string {
	var alias string
	if len(aa) > 0 {
		alias = aa[0] + "."
	}

	return []string{
		alias + "id",
		alias + "name",
		alias + "rel_module",
		alias + "place",
		alias + "kind",
		alias + "label",
		alias + "options",
		alias + "is_private",
		alias + "is_required",
		alias + "is_visible",
		alias + "is_multi",
		alias + "default_value",
		alias + "created_at",
		alias + "updated_at",
		alias + "deleted_at",
	}
}

// {false true true false}

// internalComposeModuleFieldEncoder encodes fields from types.ModuleField to store.Payload (map)
//
// Encoding is done by using generic approach or by calling encodeComposeModuleField
// func when rdbms.customEncoder=true
func (s Store) internalComposeModuleFieldEncoder(res *types.ModuleField) store.Payload {
	return store.Payload{
		"id":            res.ID,
		"name":          res.Name,
		"rel_module":    res.ModuleID,
		"place":         res.Place,
		"kind":          res.Kind,
		"label":         res.Label,
		"options":       res.Options,
		"is_private":    res.Private,
		"is_required":   res.Required,
		"is_visible":    res.Visible,
		"is_multi":      res.Multi,
		"default_value": res.DefaultValue,
		"created_at":    res.CreatedAt,
		"updated_at":    res.UpdatedAt,
		"deleted_at":    res.DeletedAt,
	}
}
