package requests

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
)

type RegisterUserRequest struct {
	Name     string `form:"name" json:"name"`
	Surname  string `form:"surname" json:"surname"`
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
}

func (r *RegisterUserRequest) Authorize(ctx http.Context) error {

	return nil
}

func (r *RegisterUserRequest) Rules(ctx http.Context) map[string]string {
	return map[string]string{
		"name":     "required|string|min_len:2|max_len:50",
		"surname":  "required|string|min_len:2|max_len:50",
		"email":    "required|email|max_len:255",
		"password": "required|string|min_len:8|max_len:255",
	}
}

func (r *RegisterUserRequest) Messages() map[string]string {
	return map[string]string{
		"name.required":     "Name is required",
		"name.min_len":      "Name must be at least 2 characters",
		"name.max_len":      "Name cannot exceed 50 characters",
		"surname.required":  "Surname is required",
		"surname.min_len":   "Surname must be at least 2 characters",
		"surname.max_len":   "Surname cannot exceed 50 characters",
		"email.required":    "Email address is required",
		"email.email":       "Please provide a valid email address",
		"email.max_len":     "Email cannot exceed 255 characters",
		"password.required": "Password is required",
		"password.min_len":  "Password must be at least 8 characters",
		"password.max_len":  "Password cannot exceed 255 characters",
	}
}

func (r *RegisterUserRequest) Attributes() map[string]string {
	return map[string]string{
		"name":     "first name",
		"surname":  "last name",
		"email":    "email address",
		"password": "password",
	}
}

func (r *RegisterUserRequest) Filters(ctx http.Context) map[string]string {
	return map[string]string{
		"name":    "trim",
		"surname": "trim",
		"email":   "trim|lower",
	}
}

func (r *RegisterUserRequest) PrepareForValidation(ctx http.Context, data validation.Data) error {

	return nil
}
