package request

// This file is auto-generated.
//
// Changes to this file may cause incorrect behavior and will be lost if
// the code is regenerated.
//
// Definitions file that controls how this file is generated:
//

import (
	"encoding/json"
	"fmt"
	"github.com/cortezaproject/corteza-server/pkg/payload"
	"github.com/go-chi/chi"
	"io"
	"mime/multipart"
	"net/http"
	"strings"
)

// dummy vars to prevent
// unused imports complain
var (
	_ = chi.URLParam
	_ = multipart.ErrMessageTooLarge
	_ = payload.ParseUint64s
)

type (
	// Internal API interface
	AttachmentList struct {
		// Kind PATH parameter
		//
		// Attachment kind
		Kind string

		// NamespaceID PATH parameter
		//
		// Namespace ID
		NamespaceID uint64 `json:",string"`

		// Sign GET parameter
		//
		// Signature
		Sign string

		// UserID GET parameter
		//
		// User ID
		UserID uint64 `json:",string"`

		// PageID GET parameter
		//
		// Filter attachments by page ID
		PageID uint64 `json:",string"`

		// ModuleID GET parameter
		//
		// Filter attachments by module ID
		ModuleID uint64 `json:",string"`

		// RecordID GET parameter
		//
		// Filter attachments by record ID
		RecordID uint64 `json:",string"`

		// FieldName GET parameter
		//
		// Filter attachments by field name
		FieldName string

		// Limit GET parameter
		//
		// Limit
		Limit uint

		// Offset GET parameter
		//
		// Offset
		Offset uint

		// Page GET parameter
		//
		// Page number (1-based)
		Page uint

		// PerPage GET parameter
		//
		// Returned items per page (default 50)
		PerPage uint
	}

	AttachmentRead struct {
		// Kind PATH parameter
		//
		// Attachment kind
		Kind string

		// NamespaceID PATH parameter
		//
		// Namespace ID
		NamespaceID uint64 `json:",string"`

		// AttachmentID PATH parameter
		//
		// Attachment ID
		AttachmentID uint64 `json:",string"`

		// Sign GET parameter
		//
		// Signature
		Sign string

		// UserID GET parameter
		//
		// User ID
		UserID uint64 `json:",string"`
	}

	AttachmentDelete struct {
		// Kind PATH parameter
		//
		// Attachment kind
		Kind string

		// NamespaceID PATH parameter
		//
		// Namespace ID
		NamespaceID uint64 `json:",string"`

		// AttachmentID PATH parameter
		//
		// Attachment ID
		AttachmentID uint64 `json:",string"`

		// Sign GET parameter
		//
		// Signature
		Sign string

		// UserID GET parameter
		//
		// User ID
		UserID uint64 `json:",string"`
	}

	AttachmentOriginal struct {
		// Kind PATH parameter
		//
		// Attachment kind
		Kind string

		// NamespaceID PATH parameter
		//
		// Namespace ID
		NamespaceID uint64 `json:",string"`

		// AttachmentID PATH parameter
		//
		// Attachment ID
		AttachmentID uint64 `json:",string"`

		// Name PATH parameter
		//
		// File name
		Name string

		// Sign GET parameter
		//
		// Signature
		Sign string

		// UserID GET parameter
		//
		// User ID
		UserID uint64 `json:",string"`

		// Download GET parameter
		//
		// Force file download
		Download bool
	}

	AttachmentPreview struct {
		// Kind PATH parameter
		//
		// Attachment kind
		Kind string

		// NamespaceID PATH parameter
		//
		// Namespace ID
		NamespaceID uint64 `json:",string"`

		// AttachmentID PATH parameter
		//
		// Attachment ID
		AttachmentID uint64 `json:",string"`

		// Ext PATH parameter
		//
		// Preview extension/format
		Ext string

		// Sign GET parameter
		//
		// Signature
		Sign string

		// UserID GET parameter
		//
		// User ID
		UserID uint64 `json:",string"`
	}
)

// NewAttachmentList request
func NewAttachmentList() *AttachmentList {
	return &AttachmentList{}
}

// Auditable returns all auditable/loggable parameters
func (r AttachmentList) Auditable() map[string]interface{} {
	return map[string]interface{}{
		"kind":        r.Kind,
		"namespaceID": r.NamespaceID,
		"sign":        r.Sign,
		"userID":      r.UserID,
		"pageID":      r.PageID,
		"moduleID":    r.ModuleID,
		"recordID":    r.RecordID,
		"fieldName":   r.FieldName,
		"limit":       r.Limit,
		"offset":      r.Offset,
		"page":        r.Page,
		"perPage":     r.PerPage,
	}
}

// Auditable returns all auditable/loggable parameters
func (r AttachmentList) GetKind() string {
	return r.Kind
}

// Auditable returns all auditable/loggable parameters
func (r AttachmentList) GetNamespaceID() uint64 {
	return r.NamespaceID
}

// Auditable returns all auditable/loggable parameters
func (r AttachmentList) GetSign() string {
	return r.Sign
}

// Auditable returns all auditable/loggable parameters
func (r AttachmentList) GetUserID() uint64 {
	return r.UserID
}

// Auditable returns all auditable/loggable parameters
func (r AttachmentList) GetPageID() uint64 {
	return r.PageID
}

// Auditable returns all auditable/loggable parameters
func (r AttachmentList) GetModuleID() uint64 {
	return r.ModuleID
}

// Auditable returns all auditable/loggable parameters
func (r AttachmentList) GetRecordID() uint64 {
	return r.RecordID
}

// Auditable returns all auditable/loggable parameters
func (r AttachmentList) GetFieldName() string {
	return r.FieldName
}

// Auditable returns all auditable/loggable parameters
func (r AttachmentList) GetLimit() uint {
	return r.Limit
}

// Auditable returns all auditable/loggable parameters
func (r AttachmentList) GetOffset() uint {
	return r.Offset
}

// Auditable returns all auditable/loggable parameters
func (r AttachmentList) GetPage() uint {
	return r.Page
}

// Auditable returns all auditable/loggable parameters
func (r AttachmentList) GetPerPage() uint {
	return r.PerPage
}

// Fill processes request and fills internal variables
func (r *AttachmentList) Fill(req *http.Request) (err error) {
	if strings.ToLower(req.Header.Get("content-type")) == "application/json" {
		err = json.NewDecoder(req.Body).Decode(r)

		switch {
		case err == io.EOF:
			err = nil
		case err != nil:
			return fmt.Errorf("error parsing http request body: %w", err)
		}
	}

	{
		// GET params
		tmp := req.URL.Query()

		if val, ok := tmp["sign"]; ok && len(val) > 0 {
			r.Sign, err = val[0], nil
			if err != nil {
				return err
			}
		}
		if val, ok := tmp["userID"]; ok && len(val) > 0 {
			r.UserID, err = payload.ParseUint64(val[0]), nil
			if err != nil {
				return err
			}
		}
		if val, ok := tmp["pageID"]; ok && len(val) > 0 {
			r.PageID, err = payload.ParseUint64(val[0]), nil
			if err != nil {
				return err
			}
		}
		if val, ok := tmp["moduleID"]; ok && len(val) > 0 {
			r.ModuleID, err = payload.ParseUint64(val[0]), nil
			if err != nil {
				return err
			}
		}
		if val, ok := tmp["recordID"]; ok && len(val) > 0 {
			r.RecordID, err = payload.ParseUint64(val[0]), nil
			if err != nil {
				return err
			}
		}
		if val, ok := tmp["fieldName"]; ok && len(val) > 0 {
			r.FieldName, err = val[0], nil
			if err != nil {
				return err
			}
		}
		if val, ok := tmp["limit"]; ok && len(val) > 0 {
			r.Limit, err = payload.ParseUint(val[0]), nil
			if err != nil {
				return err
			}
		}
		if val, ok := tmp["offset"]; ok && len(val) > 0 {
			r.Offset, err = payload.ParseUint(val[0]), nil
			if err != nil {
				return err
			}
		}
		if val, ok := tmp["page"]; ok && len(val) > 0 {
			r.Page, err = payload.ParseUint(val[0]), nil
			if err != nil {
				return err
			}
		}
		if val, ok := tmp["perPage"]; ok && len(val) > 0 {
			r.PerPage, err = payload.ParseUint(val[0]), nil
			if err != nil {
				return err
			}
		}
	}

	{
		var val string
		// path params

		val = chi.URLParam(req, "kind")
		r.Kind, err = val, nil
		if err != nil {
			return err
		}

		val = chi.URLParam(req, "namespaceID")
		r.NamespaceID, err = payload.ParseUint64(val), nil
		if err != nil {
			return err
		}

	}

	return err
}

// NewAttachmentRead request
func NewAttachmentRead() *AttachmentRead {
	return &AttachmentRead{}
}

// Auditable returns all auditable/loggable parameters
func (r AttachmentRead) Auditable() map[string]interface{} {
	return map[string]interface{}{
		"kind":         r.Kind,
		"namespaceID":  r.NamespaceID,
		"attachmentID": r.AttachmentID,
		"sign":         r.Sign,
		"userID":       r.UserID,
	}
}

// Auditable returns all auditable/loggable parameters
func (r AttachmentRead) GetKind() string {
	return r.Kind
}

// Auditable returns all auditable/loggable parameters
func (r AttachmentRead) GetNamespaceID() uint64 {
	return r.NamespaceID
}

// Auditable returns all auditable/loggable parameters
func (r AttachmentRead) GetAttachmentID() uint64 {
	return r.AttachmentID
}

// Auditable returns all auditable/loggable parameters
func (r AttachmentRead) GetSign() string {
	return r.Sign
}

// Auditable returns all auditable/loggable parameters
func (r AttachmentRead) GetUserID() uint64 {
	return r.UserID
}

// Fill processes request and fills internal variables
func (r *AttachmentRead) Fill(req *http.Request) (err error) {
	if strings.ToLower(req.Header.Get("content-type")) == "application/json" {
		err = json.NewDecoder(req.Body).Decode(r)

		switch {
		case err == io.EOF:
			err = nil
		case err != nil:
			return fmt.Errorf("error parsing http request body: %w", err)
		}
	}

	{
		// GET params
		tmp := req.URL.Query()

		if val, ok := tmp["sign"]; ok && len(val) > 0 {
			r.Sign, err = val[0], nil
			if err != nil {
				return err
			}
		}
		if val, ok := tmp["userID"]; ok && len(val) > 0 {
			r.UserID, err = payload.ParseUint64(val[0]), nil
			if err != nil {
				return err
			}
		}
	}

	{
		var val string
		// path params

		val = chi.URLParam(req, "kind")
		r.Kind, err = val, nil
		if err != nil {
			return err
		}

		val = chi.URLParam(req, "namespaceID")
		r.NamespaceID, err = payload.ParseUint64(val), nil
		if err != nil {
			return err
		}

		val = chi.URLParam(req, "attachmentID")
		r.AttachmentID, err = payload.ParseUint64(val), nil
		if err != nil {
			return err
		}

	}

	return err
}

// NewAttachmentDelete request
func NewAttachmentDelete() *AttachmentDelete {
	return &AttachmentDelete{}
}

// Auditable returns all auditable/loggable parameters
func (r AttachmentDelete) Auditable() map[string]interface{} {
	return map[string]interface{}{
		"kind":         r.Kind,
		"namespaceID":  r.NamespaceID,
		"attachmentID": r.AttachmentID,
		"sign":         r.Sign,
		"userID":       r.UserID,
	}
}

// Auditable returns all auditable/loggable parameters
func (r AttachmentDelete) GetKind() string {
	return r.Kind
}

// Auditable returns all auditable/loggable parameters
func (r AttachmentDelete) GetNamespaceID() uint64 {
	return r.NamespaceID
}

// Auditable returns all auditable/loggable parameters
func (r AttachmentDelete) GetAttachmentID() uint64 {
	return r.AttachmentID
}

// Auditable returns all auditable/loggable parameters
func (r AttachmentDelete) GetSign() string {
	return r.Sign
}

// Auditable returns all auditable/loggable parameters
func (r AttachmentDelete) GetUserID() uint64 {
	return r.UserID
}

// Fill processes request and fills internal variables
func (r *AttachmentDelete) Fill(req *http.Request) (err error) {
	if strings.ToLower(req.Header.Get("content-type")) == "application/json" {
		err = json.NewDecoder(req.Body).Decode(r)

		switch {
		case err == io.EOF:
			err = nil
		case err != nil:
			return fmt.Errorf("error parsing http request body: %w", err)
		}
	}

	{
		// GET params
		tmp := req.URL.Query()

		if val, ok := tmp["sign"]; ok && len(val) > 0 {
			r.Sign, err = val[0], nil
			if err != nil {
				return err
			}
		}
		if val, ok := tmp["userID"]; ok && len(val) > 0 {
			r.UserID, err = payload.ParseUint64(val[0]), nil
			if err != nil {
				return err
			}
		}
	}

	{
		var val string
		// path params

		val = chi.URLParam(req, "kind")
		r.Kind, err = val, nil
		if err != nil {
			return err
		}

		val = chi.URLParam(req, "namespaceID")
		r.NamespaceID, err = payload.ParseUint64(val), nil
		if err != nil {
			return err
		}

		val = chi.URLParam(req, "attachmentID")
		r.AttachmentID, err = payload.ParseUint64(val), nil
		if err != nil {
			return err
		}

	}

	return err
}

// NewAttachmentOriginal request
func NewAttachmentOriginal() *AttachmentOriginal {
	return &AttachmentOriginal{}
}

// Auditable returns all auditable/loggable parameters
func (r AttachmentOriginal) Auditable() map[string]interface{} {
	return map[string]interface{}{
		"kind":         r.Kind,
		"namespaceID":  r.NamespaceID,
		"attachmentID": r.AttachmentID,
		"name":         r.Name,
		"sign":         r.Sign,
		"userID":       r.UserID,
		"download":     r.Download,
	}
}

// Auditable returns all auditable/loggable parameters
func (r AttachmentOriginal) GetKind() string {
	return r.Kind
}

// Auditable returns all auditable/loggable parameters
func (r AttachmentOriginal) GetNamespaceID() uint64 {
	return r.NamespaceID
}

// Auditable returns all auditable/loggable parameters
func (r AttachmentOriginal) GetAttachmentID() uint64 {
	return r.AttachmentID
}

// Auditable returns all auditable/loggable parameters
func (r AttachmentOriginal) GetName() string {
	return r.Name
}

// Auditable returns all auditable/loggable parameters
func (r AttachmentOriginal) GetSign() string {
	return r.Sign
}

// Auditable returns all auditable/loggable parameters
func (r AttachmentOriginal) GetUserID() uint64 {
	return r.UserID
}

// Auditable returns all auditable/loggable parameters
func (r AttachmentOriginal) GetDownload() bool {
	return r.Download
}

// Fill processes request and fills internal variables
func (r *AttachmentOriginal) Fill(req *http.Request) (err error) {
	if strings.ToLower(req.Header.Get("content-type")) == "application/json" {
		err = json.NewDecoder(req.Body).Decode(r)

		switch {
		case err == io.EOF:
			err = nil
		case err != nil:
			return fmt.Errorf("error parsing http request body: %w", err)
		}
	}

	{
		// GET params
		tmp := req.URL.Query()

		if val, ok := tmp["sign"]; ok && len(val) > 0 {
			r.Sign, err = val[0], nil
			if err != nil {
				return err
			}
		}
		if val, ok := tmp["userID"]; ok && len(val) > 0 {
			r.UserID, err = payload.ParseUint64(val[0]), nil
			if err != nil {
				return err
			}
		}
		if val, ok := tmp["download"]; ok && len(val) > 0 {
			r.Download, err = payload.ParseBool(val[0]), nil
			if err != nil {
				return err
			}
		}
	}

	{
		var val string
		// path params

		val = chi.URLParam(req, "kind")
		r.Kind, err = val, nil
		if err != nil {
			return err
		}

		val = chi.URLParam(req, "namespaceID")
		r.NamespaceID, err = payload.ParseUint64(val), nil
		if err != nil {
			return err
		}

		val = chi.URLParam(req, "attachmentID")
		r.AttachmentID, err = payload.ParseUint64(val), nil
		if err != nil {
			return err
		}

		val = chi.URLParam(req, "name")
		r.Name, err = val, nil
		if err != nil {
			return err
		}

	}

	return err
}

// NewAttachmentPreview request
func NewAttachmentPreview() *AttachmentPreview {
	return &AttachmentPreview{}
}

// Auditable returns all auditable/loggable parameters
func (r AttachmentPreview) Auditable() map[string]interface{} {
	return map[string]interface{}{
		"kind":         r.Kind,
		"namespaceID":  r.NamespaceID,
		"attachmentID": r.AttachmentID,
		"ext":          r.Ext,
		"sign":         r.Sign,
		"userID":       r.UserID,
	}
}

// Auditable returns all auditable/loggable parameters
func (r AttachmentPreview) GetKind() string {
	return r.Kind
}

// Auditable returns all auditable/loggable parameters
func (r AttachmentPreview) GetNamespaceID() uint64 {
	return r.NamespaceID
}

// Auditable returns all auditable/loggable parameters
func (r AttachmentPreview) GetAttachmentID() uint64 {
	return r.AttachmentID
}

// Auditable returns all auditable/loggable parameters
func (r AttachmentPreview) GetExt() string {
	return r.Ext
}

// Auditable returns all auditable/loggable parameters
func (r AttachmentPreview) GetSign() string {
	return r.Sign
}

// Auditable returns all auditable/loggable parameters
func (r AttachmentPreview) GetUserID() uint64 {
	return r.UserID
}

// Fill processes request and fills internal variables
func (r *AttachmentPreview) Fill(req *http.Request) (err error) {
	if strings.ToLower(req.Header.Get("content-type")) == "application/json" {
		err = json.NewDecoder(req.Body).Decode(r)

		switch {
		case err == io.EOF:
			err = nil
		case err != nil:
			return fmt.Errorf("error parsing http request body: %w", err)
		}
	}

	{
		// GET params
		tmp := req.URL.Query()

		if val, ok := tmp["sign"]; ok && len(val) > 0 {
			r.Sign, err = val[0], nil
			if err != nil {
				return err
			}
		}
		if val, ok := tmp["userID"]; ok && len(val) > 0 {
			r.UserID, err = payload.ParseUint64(val[0]), nil
			if err != nil {
				return err
			}
		}
	}

	{
		var val string
		// path params

		val = chi.URLParam(req, "kind")
		r.Kind, err = val, nil
		if err != nil {
			return err
		}

		val = chi.URLParam(req, "namespaceID")
		r.NamespaceID, err = payload.ParseUint64(val), nil
		if err != nil {
			return err
		}

		val = chi.URLParam(req, "attachmentID")
		r.AttachmentID, err = payload.ParseUint64(val), nil
		if err != nil {
			return err
		}

		val = chi.URLParam(req, "ext")
		r.Ext, err = val, nil
		if err != nil {
			return err
		}

	}

	return err
}
