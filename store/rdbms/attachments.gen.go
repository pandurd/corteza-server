package rdbms

// This file is an auto-generated file
//
// Template:    pkg/codegen/assets/store_rdbms.gen.go.tpl
// Definitions: store/attachments.yaml
//
// Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated.

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/cortezaproject/corteza-server/store"
	"github.com/cortezaproject/corteza-server/system/types"
)

// SearchAttachments returns all matching rows
//
// This function calls convertAttachmentFilter with the given
// types.AttachmentFilter and expects to receive a working squirrel.SelectBuilder
func (s Store) SearchAttachments(ctx context.Context, f types.AttachmentFilter) (types.AttachmentSet, types.AttachmentFilter, error) {
	var scap uint
	q, err := s.convertAttachmentFilter(f)
	if err != nil {
		return nil, f, err
	}

	if scap == 0 {
		scap = DefaultSliceCapacity
	}

	var (
		set = make([]*types.Attachment, 0, scap)
		// Paging is disabled in definition yaml file
		// {search: {disablePaging:true}} and this allows
		// a much simpler row fetching logic
		fetch = func() error {
			var (
				res       *types.Attachment
				rows, err = s.Query(ctx, q)
			)

			if err != nil {
				return err
			}

			for rows.Next() {
				if res, err = s.internalAttachmentRowScanner(rows, rows.Err()); err != nil {
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

// LookupAttachmentByID searches for attachment by its ID
//
// It returns attachment even if deleted
func (s Store) LookupAttachmentByID(ctx context.Context, id uint64) (*types.Attachment, error) {
	return s.AttachmentLookup(ctx, squirrel.Eq{
		"att.id": id,
	})
}

// CreateAttachment creates one or more rows in attachments table
func (s Store) CreateAttachment(ctx context.Context, rr ...*types.Attachment) (err error) {
	for _, res := range rr {
		err = s.Exec(ctx, s.InsertBuilder(s.AttachmentTable()).SetMap(s.internalAttachmentEncoder(res)))
		if err != nil {
			return s.config.ErrorHandler(err)
		}
	}

	return
}

// UpdateAttachment updates one or more existing rows in attachments
func (s Store) UpdateAttachment(ctx context.Context, rr ...*types.Attachment) error {
	return s.config.ErrorHandler(s.PartialAttachmentUpdate(ctx, nil, rr...))
}

// PartialAttachmentUpdate updates one or more existing rows in attachments
func (s Store) PartialAttachmentUpdate(ctx context.Context, onlyColumns []string, rr ...*types.Attachment) (err error) {
	for _, res := range rr {
		err = s.ExecUpdateAttachments(
			ctx,
			squirrel.Eq{s.preprocessColumn("att.id", ""): s.preprocessValue(res.ID, "")},
			s.internalAttachmentEncoder(res).Skip("id").Only(onlyColumns...))
		if err != nil {
			return s.config.ErrorHandler(err)
		}
	}

	return
}

// RemoveAttachment removes one or more rows from attachments table
func (s Store) RemoveAttachment(ctx context.Context, rr ...*types.Attachment) (err error) {
	for _, res := range rr {
		err = s.Exec(ctx, s.DeleteBuilder(s.AttachmentTable("att")).Where(squirrel.Eq{s.preprocessColumn("att.id", ""): s.preprocessValue(res.ID, "")}))
		if err != nil {
			return s.config.ErrorHandler(err)
		}
	}

	return nil
}

// RemoveAttachmentByID removes row from the attachments table
func (s Store) RemoveAttachmentByID(ctx context.Context, ID uint64) error {
	return s.config.ErrorHandler(s.Exec(ctx, s.DeleteBuilder(s.AttachmentTable("att")).Where(squirrel.Eq{s.preprocessColumn("att.id", ""): s.preprocessValue(ID, "")})))
}

// TruncateAttachments removes all rows from the attachments table
func (s Store) TruncateAttachments(ctx context.Context) error {
	return s.config.ErrorHandler(s.Truncate(ctx, s.AttachmentTable()))
}

// ExecUpdateAttachments updates all matched (by cnd) rows in attachments with given data
func (s Store) ExecUpdateAttachments(ctx context.Context, cnd squirrel.Sqlizer, set store.Payload) error {
	return s.config.ErrorHandler(s.Exec(ctx, s.UpdateBuilder(s.AttachmentTable("att")).Where(cnd).SetMap(set)))
}

// AttachmentLookup prepares Attachment query and executes it,
// returning types.Attachment (or error)
func (s Store) AttachmentLookup(ctx context.Context, cnd squirrel.Sqlizer) (*types.Attachment, error) {
	return s.internalAttachmentRowScanner(s.QueryRow(ctx, s.QueryAttachments().Where(cnd)))
}

func (s Store) internalAttachmentRowScanner(row rowScanner, err error) (*types.Attachment, error) {
	if err != nil {
		return nil, err
	}

	var res = &types.Attachment{}
	if _, has := s.config.RowScanners["attachment"]; has {
		scanner := s.config.RowScanners["attachment"].(func(rowScanner, *types.Attachment) error)
		err = scanner(row, res)
	} else {
		err = row.Scan(
			&res.ID,
			&res.OwnerID,
			&res.Kind,
			&res.Url,
			&res.PreviewUrl,
			&res.Name,
			&res.Meta,
			&res.CreatedAt,
			&res.UpdatedAt,
			&res.DeletedAt,
		)
	}

	if err == sql.ErrNoRows {
		return nil, store.ErrNotFound
	}

	if err != nil {
		return nil, fmt.Errorf("could not scan db row for Attachment: %w", err)
	} else {
		return res, nil
	}
}

// QueryAttachments returns squirrel.SelectBuilder with set table and all columns
func (s Store) QueryAttachments() squirrel.SelectBuilder {
	return s.SelectBuilder(s.AttachmentTable("att"), s.AttachmentColumns("att")...)
}

// AttachmentTable name of the db table
func (Store) AttachmentTable(aa ...string) string {
	var alias string
	if len(aa) > 0 {
		alias = " AS " + aa[0]
	}

	return "attachments" + alias
}

// AttachmentColumns returns all defined table columns
//
// With optional string arg, all columns are returned aliased
func (Store) AttachmentColumns(aa ...string) []string {
	var alias string
	if len(aa) > 0 {
		alias = aa[0] + "."
	}

	return []string{
		alias + "id",
		alias + "rel_owner",
		alias + "kind",
		alias + "url",
		alias + "preview_url",
		alias + "name",
		alias + "meta",
		alias + "created_at",
		alias + "updated_at",
		alias + "deleted_at",
	}
}

// {false true true false}

// internalAttachmentEncoder encodes fields from types.Attachment to store.Payload (map)
//
// Encoding is done by using generic approach or by calling encodeAttachment
// func when rdbms.customEncoder=true
func (s Store) internalAttachmentEncoder(res *types.Attachment) store.Payload {
	return store.Payload{
		"id":          res.ID,
		"rel_owner":   res.OwnerID,
		"kind":        res.Kind,
		"url":         res.Url,
		"preview_url": res.PreviewUrl,
		"name":        res.Name,
		"meta":        res.Meta,
		"created_at":  res.CreatedAt,
		"updated_at":  res.UpdatedAt,
		"deleted_at":  res.DeletedAt,
	}
}