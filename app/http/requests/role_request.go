package requests

import (
	"github.com/goravel/framework/contracts/http"
)

type RoleRequest struct {
	Role        string `form:"role" json:"role"`
	Description string `form:"description" json:"description"`
	IsActive    string `form:"is_active" json:"is_active"`
}

func (r *RoleRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *RoleRequest) Rules(ctx http.Context) map[string]string {
	return map[string]string{
		"role":        "required|in:super_admin,admin,manager,partner,client",
		"description": "required|string",
		"is_active":   "required|bool",
	}
}
