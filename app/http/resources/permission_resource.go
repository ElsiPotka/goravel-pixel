package resources

import (
	"pixel/app/models"
	"time"

	"github.com/google/uuid"
)

type PermissionResource struct {
	ID          uuid.UUID            `json:"id"`
	Permission  string               `json:"permission"`
	Description string               `json:"description"`
	IsActive    bool                 `json:"is_active"`
	Roles       []RoleResourceSimple `json:"roles,omitempty"`
	CreatedAt   time.Time            `json:"created_at"`
	UpdatedAt   time.Time            `json:"updated_at"`
}

type PermissionResourceSimple struct {
	ID          uuid.UUID `json:"id"`
	Permission  string    `json:"permission"`
	Description string    `json:"description"`
	IsActive    bool      `json:"is_active"`
}

func NewPermissionResource(permission *models.Permission) *PermissionResource {
	permissionResource := &PermissionResource{
		ID:          permission.ID,
		Permission:  permission.Permission,
		Description: permission.Description,
		IsActive:    permission.IsActive,
		CreatedAt:   permission.CreatedAt,
		UpdatedAt:   permission.UpdatedAt,
	}

	if len(permission.Roles) > 0 {
		permissionResource.Roles = make([]RoleResourceSimple, len(permission.Roles))
		for i, role := range permission.Roles {
			permissionResource.Roles[i] = *NewRoleResourceSimple(&role)
		}
	}

	return permissionResource
}

func NewPermissionResourceSimple(permission *models.Permission) *PermissionResourceSimple {
	return &PermissionResourceSimple{
		ID:          permission.ID,
		Permission:  permission.Permission,
		Description: permission.Description,
		IsActive:    permission.IsActive,
	}
}

func NewPermissionResourceCollection(permissions []models.Permission) []*PermissionResource {
	resources := make([]*PermissionResource, len(permissions))
	for i, permission := range permissions {
		resources[i] = NewPermissionResource(&permission)
	}
	return resources
}

func NewPermissionResourceSimpleCollection(permissions []models.Permission) []*PermissionResourceSimple {
	resources := make([]*PermissionResourceSimple, len(permissions))
	for i, permission := range permissions {
		resources[i] = NewPermissionResourceSimple(&permission)
	}
	return resources
}
