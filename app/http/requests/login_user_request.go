package requests

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type LoginUserRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

func (r *LoginUserRequest) Authorize(ctx http.Context) error {

	return nil
}

func (r *LoginUserRequest) Rules(ctx http.Context) map[string]string {
	return map[string]string{
		"email":    "required|email",
		"password": "required|string",
	}
}

func (r *LoginUserRequest) Messages() map[string]string {
	return map[string]string{
		"email.required":    "Email address is required",
		"email.email":       "Please provide a valid email address",
		"password.required": "Password is required",
	}
}

func (r *LoginUserRequest) Attributes() map[string]string {
	return map[string]string{
		"email":    "email address",
		"password": "password",
	}
}

func (r *LoginUserRequest) Filters(ctx http.Context) map[string]string {
	return map[string]string{
		"email": "trim|lower",
	}
}

func (r *LoginUserRequest) PrepareForValidation(ctx http.Context, data validation.Data) error {

	return nil
}
