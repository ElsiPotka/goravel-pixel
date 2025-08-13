package resources

import (
	"pixel/app/models"
	"time"

	"github.com/google/uuid"
)

type RoleResource struct {
	ID          uuid.UUID                  `json:"id"`
	Role        models.RoleType            `json:"role"`
	Description string                     `json:"description"`
	IsActive    bool                       `json:"is_active"`
	Permissions []PermissionResourceSimple `json:"permissions,omitempty"`
	CreatedAt   time.Time                  `json:"created_at"`
	UpdatedAt   time.Time                  `json:"updated_at"`
}

type RoleResourceSimple struct {
	ID          uuid.UUID       `json:"id"`
	Role        models.RoleType `json:"role"`
	Description string          `json:"description"`
	IsActive    bool            `json:"is_active"`
}

func NewRoleResource(role *models.Role) *RoleResource {
	roleResource := &RoleResource{
		ID:          role.ID,
		Role:        role.Role,
		Description: role.Description,
		IsActive:    role.IsActive,
		CreatedAt:   role.CreatedAt,
		UpdatedAt:   role.UpdatedAt,
	}

	if len(role.Permissions) > 0 {
		roleResource.Permissions = make([]PermissionResourceSimple, len(role.Permissions))
		for i, permission := range role.Permissions {
			roleResource.Permissions[i] = *NewPermissionResourceSimple(&permission)
		}
	}

	return roleResource
}

func NewRoleResourceSimple(role *models.Role) *RoleResourceSimple {
	return &RoleResourceSimple{
		ID:          role.ID,
		Role:        role.Role,
		Description: role.Description,
		IsActive:    role.IsActive,
	}
}

func NewRoleResourceCollection(roles []models.Role) []*RoleResource {
	resources := make([]*RoleResource, len(roles))
	for i, role := range roles {
		resources[i] = NewRoleResource(&role)
	}
	return resources
}

func NewRoleResourceSimpleCollection(roles []models.Role) []*RoleResourceSimple {
	resources := make([]*RoleResourceSimple, len(roles))
	for i := range roles {
		resources[i] = NewRoleResourceSimple(&roles[i])
	}
	return resources
}
