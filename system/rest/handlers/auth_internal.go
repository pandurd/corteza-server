package handlers

// This file is auto-generated.
//
// Changes to this file may cause incorrect behavior and will be lost if
// the code is regenerated.
//
// Definitions file that controls how this file is generated:
//

import (
	"context"
	"github.com/go-chi/chi"
	"github.com/titpetric/factory/resputil"
	"net/http"

	"github.com/cortezaproject/corteza-server/pkg/logger"
	"github.com/cortezaproject/corteza-server/system/rest/request"
)

type (
	// Internal API interface
	Auth_internalAPI interface {
		Login(context.Context, *request.Auth_internalLogin) (interface{}, error)
		Signup(context.Context, *request.Auth_internalSignup) (interface{}, error)
		RequestPasswordReset(context.Context, *request.Auth_internalRequestPasswordReset) (interface{}, error)
		ExchangePasswordResetToken(context.Context, *request.Auth_internalExchangePasswordResetToken) (interface{}, error)
		ResetPassword(context.Context, *request.Auth_internalResetPassword) (interface{}, error)
		ConfirmEmail(context.Context, *request.Auth_internalConfirmEmail) (interface{}, error)
		ChangePassword(context.Context, *request.Auth_internalChangePassword) (interface{}, error)
	}

	// HTTP API interface
	Auth_internal struct {
		Login                      func(http.ResponseWriter, *http.Request)
		Signup                     func(http.ResponseWriter, *http.Request)
		RequestPasswordReset       func(http.ResponseWriter, *http.Request)
		ExchangePasswordResetToken func(http.ResponseWriter, *http.Request)
		ResetPassword              func(http.ResponseWriter, *http.Request)
		ConfirmEmail               func(http.ResponseWriter, *http.Request)
		ChangePassword             func(http.ResponseWriter, *http.Request)
	}
)

func NewAuth_internal(h Auth_internalAPI) *Auth_internal {
	return &Auth_internal{
		Login: func(w http.ResponseWriter, r *http.Request) {
			defer r.Body.Close()
			params := request.NewAuth_internalLogin()
			if err := params.Fill(r); err != nil {
				logger.LogParamError("Auth_internal.Login", r, err)
				resputil.JSON(w, err)
				return
			}

			value, err := h.Login(r.Context(), params)
			if err != nil {
				logger.LogControllerError("Auth_internal.Login", r, err, params.Auditable())
				resputil.JSON(w, err)
				return
			}
			logger.LogControllerCall("Auth_internal.Login", r, params.Auditable())
			if !serveHTTP(value, w, r) {
				resputil.JSON(w, value)
			}
		},
		Signup: func(w http.ResponseWriter, r *http.Request) {
			defer r.Body.Close()
			params := request.NewAuth_internalSignup()
			if err := params.Fill(r); err != nil {
				logger.LogParamError("Auth_internal.Signup", r, err)
				resputil.JSON(w, err)
				return
			}

			value, err := h.Signup(r.Context(), params)
			if err != nil {
				logger.LogControllerError("Auth_internal.Signup", r, err, params.Auditable())
				resputil.JSON(w, err)
				return
			}
			logger.LogControllerCall("Auth_internal.Signup", r, params.Auditable())
			if !serveHTTP(value, w, r) {
				resputil.JSON(w, value)
			}
		},
		RequestPasswordReset: func(w http.ResponseWriter, r *http.Request) {
			defer r.Body.Close()
			params := request.NewAuth_internalRequestPasswordReset()
			if err := params.Fill(r); err != nil {
				logger.LogParamError("Auth_internal.RequestPasswordReset", r, err)
				resputil.JSON(w, err)
				return
			}

			value, err := h.RequestPasswordReset(r.Context(), params)
			if err != nil {
				logger.LogControllerError("Auth_internal.RequestPasswordReset", r, err, params.Auditable())
				resputil.JSON(w, err)
				return
			}
			logger.LogControllerCall("Auth_internal.RequestPasswordReset", r, params.Auditable())
			if !serveHTTP(value, w, r) {
				resputil.JSON(w, value)
			}
		},
		ExchangePasswordResetToken: func(w http.ResponseWriter, r *http.Request) {
			defer r.Body.Close()
			params := request.NewAuth_internalExchangePasswordResetToken()
			if err := params.Fill(r); err != nil {
				logger.LogParamError("Auth_internal.ExchangePasswordResetToken", r, err)
				resputil.JSON(w, err)
				return
			}

			value, err := h.ExchangePasswordResetToken(r.Context(), params)
			if err != nil {
				logger.LogControllerError("Auth_internal.ExchangePasswordResetToken", r, err, params.Auditable())
				resputil.JSON(w, err)
				return
			}
			logger.LogControllerCall("Auth_internal.ExchangePasswordResetToken", r, params.Auditable())
			if !serveHTTP(value, w, r) {
				resputil.JSON(w, value)
			}
		},
		ResetPassword: func(w http.ResponseWriter, r *http.Request) {
			defer r.Body.Close()
			params := request.NewAuth_internalResetPassword()
			if err := params.Fill(r); err != nil {
				logger.LogParamError("Auth_internal.ResetPassword", r, err)
				resputil.JSON(w, err)
				return
			}

			value, err := h.ResetPassword(r.Context(), params)
			if err != nil {
				logger.LogControllerError("Auth_internal.ResetPassword", r, err, params.Auditable())
				resputil.JSON(w, err)
				return
			}
			logger.LogControllerCall("Auth_internal.ResetPassword", r, params.Auditable())
			if !serveHTTP(value, w, r) {
				resputil.JSON(w, value)
			}
		},
		ConfirmEmail: func(w http.ResponseWriter, r *http.Request) {
			defer r.Body.Close()
			params := request.NewAuth_internalConfirmEmail()
			if err := params.Fill(r); err != nil {
				logger.LogParamError("Auth_internal.ConfirmEmail", r, err)
				resputil.JSON(w, err)
				return
			}

			value, err := h.ConfirmEmail(r.Context(), params)
			if err != nil {
				logger.LogControllerError("Auth_internal.ConfirmEmail", r, err, params.Auditable())
				resputil.JSON(w, err)
				return
			}
			logger.LogControllerCall("Auth_internal.ConfirmEmail", r, params.Auditable())
			if !serveHTTP(value, w, r) {
				resputil.JSON(w, value)
			}
		},
		ChangePassword: func(w http.ResponseWriter, r *http.Request) {
			defer r.Body.Close()
			params := request.NewAuth_internalChangePassword()
			if err := params.Fill(r); err != nil {
				logger.LogParamError("Auth_internal.ChangePassword", r, err)
				resputil.JSON(w, err)
				return
			}

			value, err := h.ChangePassword(r.Context(), params)
			if err != nil {
				logger.LogControllerError("Auth_internal.ChangePassword", r, err, params.Auditable())
				resputil.JSON(w, err)
				return
			}
			logger.LogControllerCall("Auth_internal.ChangePassword", r, params.Auditable())
			if !serveHTTP(value, w, r) {
				resputil.JSON(w, value)
			}
		},
	}
}

func (h Auth_internal) MountRoutes(r chi.Router, middlewares ...func(http.Handler) http.Handler) {
	r.Group(func(r chi.Router) {
		r.Use(middlewares...)
		r.Post("/auth/internal/login", h.Login)
		r.Post("/auth/internal/signup", h.Signup)
		r.Post("/auth/internal/request-password-reset", h.RequestPasswordReset)
		r.Post("/auth/internal/exchange-password-reset-token", h.ExchangePasswordResetToken)
		r.Post("/auth/internal/reset-password", h.ResetPassword)
		r.Post("/auth/internal/confirm-email", h.ConfirmEmail)
		r.Post("/auth/internal/change-password", h.ChangePassword)
	})
}
