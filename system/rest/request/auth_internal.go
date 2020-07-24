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
	Auth_internalLogin struct {
		// Email POST parameter
		//
		// Email
		Email string

		// Password POST parameter
		//
		// Password
		Password string
	}

	Auth_internalSignup struct {
		// Email POST parameter
		//
		// Email
		Email string

		// Username POST parameter
		//
		// Username
		Username string

		// Password POST parameter
		//
		// Password
		Password string

		// Handle POST parameter
		//
		// User handle
		Handle string

		// Name POST parameter
		//
		// Display name
		Name string
	}

	Auth_internalRequestPasswordReset struct {
		// Email POST parameter
		//
		// Email
		Email string
	}

	Auth_internalExchangePasswordResetToken struct {
		// Token POST parameter
		//
		// Token
		Token string
	}

	Auth_internalResetPassword struct {
		// Token POST parameter
		//
		// Token
		Token string

		// Password POST parameter
		//
		// Password
		Password string
	}

	Auth_internalConfirmEmail struct {
		// Token POST parameter
		//
		// Token
		Token string
	}

	Auth_internalChangePassword struct {
		// OldPassword POST parameter
		//
		// Old password
		OldPassword string

		// NewPassword POST parameter
		//
		// New password
		NewPassword string
	}
)

// NewAuth_internalLogin request
func NewAuth_internalLogin() *Auth_internalLogin {
	return &Auth_internalLogin{}
}

// Auditable returns all auditable/loggable parameters
func (r Auth_internalLogin) Auditable() map[string]interface{} {
	return map[string]interface{}{
		"email":    r.Email,
		"password": r.Password,
	}
}

// Fill processes request and fills internal variables
func (r *Auth_internalLogin) Fill(req *http.Request) (err error) {
	if strings.ToLower(req.Header.Get("content-type")) == "application/json" {
		err = json.NewDecoder(req.Body).Decode(r)

		switch {
		case err == io.EOF:
			err = nil
		case err != nil:
			return fmt.Errorf("error parsing http request body: %w")
		}
	}

	{
		if err = req.ParseForm(); err != nil {
			return err
		}

		// POST params

		if val, ok := req.Form["email"]; ok && len(val) > 0 {
			r.Email, err = val[0], nil
			if err != nil {
				return err
			}
		}

		if val, ok := req.Form["password"]; ok && len(val) > 0 {
			r.Password, err = val[0], nil
			if err != nil {
				return err
			}
		}
	}

	return err
}

// NewAuth_internalSignup request
func NewAuth_internalSignup() *Auth_internalSignup {
	return &Auth_internalSignup{}
}

// Auditable returns all auditable/loggable parameters
func (r Auth_internalSignup) Auditable() map[string]interface{} {
	return map[string]interface{}{
		"email":    r.Email,
		"username": r.Username,
		"password": r.Password,
		"handle":   r.Handle,
		"name":     r.Name,
	}
}

// Fill processes request and fills internal variables
func (r *Auth_internalSignup) Fill(req *http.Request) (err error) {
	if strings.ToLower(req.Header.Get("content-type")) == "application/json" {
		err = json.NewDecoder(req.Body).Decode(r)

		switch {
		case err == io.EOF:
			err = nil
		case err != nil:
			return fmt.Errorf("error parsing http request body: %w")
		}
	}

	{
		if err = req.ParseForm(); err != nil {
			return err
		}

		// POST params

		if val, ok := req.Form["email"]; ok && len(val) > 0 {
			r.Email, err = val[0], nil
			if err != nil {
				return err
			}
		}

		if val, ok := req.Form["username"]; ok && len(val) > 0 {
			r.Username, err = val[0], nil
			if err != nil {
				return err
			}
		}

		if val, ok := req.Form["password"]; ok && len(val) > 0 {
			r.Password, err = val[0], nil
			if err != nil {
				return err
			}
		}

		if val, ok := req.Form["handle"]; ok && len(val) > 0 {
			r.Handle, err = val[0], nil
			if err != nil {
				return err
			}
		}

		if val, ok := req.Form["name"]; ok && len(val) > 0 {
			r.Name, err = val[0], nil
			if err != nil {
				return err
			}
		}
	}

	return err
}

// NewAuth_internalRequestPasswordReset request
func NewAuth_internalRequestPasswordReset() *Auth_internalRequestPasswordReset {
	return &Auth_internalRequestPasswordReset{}
}

// Auditable returns all auditable/loggable parameters
func (r Auth_internalRequestPasswordReset) Auditable() map[string]interface{} {
	return map[string]interface{}{
		"email": r.Email,
	}
}

// Fill processes request and fills internal variables
func (r *Auth_internalRequestPasswordReset) Fill(req *http.Request) (err error) {
	if strings.ToLower(req.Header.Get("content-type")) == "application/json" {
		err = json.NewDecoder(req.Body).Decode(r)

		switch {
		case err == io.EOF:
			err = nil
		case err != nil:
			return fmt.Errorf("error parsing http request body: %w")
		}
	}

	{
		if err = req.ParseForm(); err != nil {
			return err
		}

		// POST params

		if val, ok := req.Form["email"]; ok && len(val) > 0 {
			r.Email, err = val[0], nil
			if err != nil {
				return err
			}
		}
	}

	return err
}

// NewAuth_internalExchangePasswordResetToken request
func NewAuth_internalExchangePasswordResetToken() *Auth_internalExchangePasswordResetToken {
	return &Auth_internalExchangePasswordResetToken{}
}

// Auditable returns all auditable/loggable parameters
func (r Auth_internalExchangePasswordResetToken) Auditable() map[string]interface{} {
	return map[string]interface{}{
		"token": r.Token,
	}
}

// Fill processes request and fills internal variables
func (r *Auth_internalExchangePasswordResetToken) Fill(req *http.Request) (err error) {
	if strings.ToLower(req.Header.Get("content-type")) == "application/json" {
		err = json.NewDecoder(req.Body).Decode(r)

		switch {
		case err == io.EOF:
			err = nil
		case err != nil:
			return fmt.Errorf("error parsing http request body: %w")
		}
	}

	{
		if err = req.ParseForm(); err != nil {
			return err
		}

		// POST params

		if val, ok := req.Form["token"]; ok && len(val) > 0 {
			r.Token, err = val[0], nil
			if err != nil {
				return err
			}
		}
	}

	return err
}

// NewAuth_internalResetPassword request
func NewAuth_internalResetPassword() *Auth_internalResetPassword {
	return &Auth_internalResetPassword{}
}

// Auditable returns all auditable/loggable parameters
func (r Auth_internalResetPassword) Auditable() map[string]interface{} {
	return map[string]interface{}{
		"token":    r.Token,
		"password": r.Password,
	}
}

// Fill processes request and fills internal variables
func (r *Auth_internalResetPassword) Fill(req *http.Request) (err error) {
	if strings.ToLower(req.Header.Get("content-type")) == "application/json" {
		err = json.NewDecoder(req.Body).Decode(r)

		switch {
		case err == io.EOF:
			err = nil
		case err != nil:
			return fmt.Errorf("error parsing http request body: %w")
		}
	}

	{
		if err = req.ParseForm(); err != nil {
			return err
		}

		// POST params

		if val, ok := req.Form["token"]; ok && len(val) > 0 {
			r.Token, err = val[0], nil
			if err != nil {
				return err
			}
		}

		if val, ok := req.Form["password"]; ok && len(val) > 0 {
			r.Password, err = val[0], nil
			if err != nil {
				return err
			}
		}
	}

	return err
}

// NewAuth_internalConfirmEmail request
func NewAuth_internalConfirmEmail() *Auth_internalConfirmEmail {
	return &Auth_internalConfirmEmail{}
}

// Auditable returns all auditable/loggable parameters
func (r Auth_internalConfirmEmail) Auditable() map[string]interface{} {
	return map[string]interface{}{
		"token": r.Token,
	}
}

// Fill processes request and fills internal variables
func (r *Auth_internalConfirmEmail) Fill(req *http.Request) (err error) {
	if strings.ToLower(req.Header.Get("content-type")) == "application/json" {
		err = json.NewDecoder(req.Body).Decode(r)

		switch {
		case err == io.EOF:
			err = nil
		case err != nil:
			return fmt.Errorf("error parsing http request body: %w")
		}
	}

	{
		if err = req.ParseForm(); err != nil {
			return err
		}

		// POST params

		if val, ok := req.Form["token"]; ok && len(val) > 0 {
			r.Token, err = val[0], nil
			if err != nil {
				return err
			}
		}
	}

	return err
}

// NewAuth_internalChangePassword request
func NewAuth_internalChangePassword() *Auth_internalChangePassword {
	return &Auth_internalChangePassword{}
}

// Auditable returns all auditable/loggable parameters
func (r Auth_internalChangePassword) Auditable() map[string]interface{} {
	return map[string]interface{}{
		"oldPassword": r.OldPassword,
		"newPassword": r.NewPassword,
	}
}

// Fill processes request and fills internal variables
func (r *Auth_internalChangePassword) Fill(req *http.Request) (err error) {
	if strings.ToLower(req.Header.Get("content-type")) == "application/json" {
		err = json.NewDecoder(req.Body).Decode(r)

		switch {
		case err == io.EOF:
			err = nil
		case err != nil:
			return fmt.Errorf("error parsing http request body: %w")
		}
	}

	{
		if err = req.ParseForm(); err != nil {
			return err
		}

		// POST params

		if val, ok := req.Form["oldPassword"]; ok && len(val) > 0 {
			r.OldPassword, err = val[0], nil
			if err != nil {
				return err
			}
		}

		if val, ok := req.Form["newPassword"]; ok && len(val) > 0 {
			r.NewPassword, err = val[0], nil
			if err != nil {
				return err
			}
		}
	}

	return err
}
