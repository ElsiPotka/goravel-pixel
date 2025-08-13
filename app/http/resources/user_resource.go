package resources

import (
	"pixel/app/models"
	"time"

	"github.com/google/uuid"
)

type UserResource struct {
	ID        uuid.UUID            `json:"id"`
	Name      string               `json:"name"`
	Surname   string               `json:"surname"`
	Email     string               `json:"email"`
	AvatarURL string               `json:"avatar_url"`
	LastLogin *time.Time           `json:"last_login"`
	Roles     []RoleResourceSimple `json:"roles,omitempty"`
	CreatedAt time.Time            `json:"created_at"`
	UpdatedAt time.Time            `json:"updated_at"`
}

func NewUserResource(user *models.User) *UserResource {
	userResource := &UserResource{
		ID:        user.ID,
		Name:      user.Name,
		Surname:   user.Surname,
		Email:     user.Email,
		AvatarURL: user.AvatarURL,
		LastLogin: user.LastLogin,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	if len(user.Roles) > 0 {
		userResource.Roles = make([]RoleResourceSimple, len(user.Roles))
		for i, role := range user.Roles {
			userResource.Roles[i] = *NewRoleResourceSimple(&role)
		}
	}

	return userResource
}
