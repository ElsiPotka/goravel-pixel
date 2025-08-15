package services

import (
	"fmt"
	"pixel/app/models"

	"github.com/google/uuid"
	"github.com/goravel/framework/contracts/database/orm"

	"github.com/goravel/framework/errors"
	"github.com/goravel/framework/facades"
)

type UserService struct {
	roleService *RoleService
}

func NewUserService() *UserService {
	return &UserService{
		roleService: NewRoleService(),
	}
}

func (s *UserService) IsEmailTaken(email string) (bool, error) {
	return facades.Orm().
		Query().
		Model(&models.User{}).
		Where("email", email).
		Exists()
}

func (s *UserService) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := facades.Orm().
		Query().
		Model(&models.User{}).
		Where("email", email).
		FirstOrFail(&user)

	if err != nil {
		if errors.Is(err, errors.OrmRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (s *UserService) GetUserWithSocialAccounts(email string) (*models.User, error) {
	var user models.User
	err := facades.Orm().
		Query().
		Model(&models.User{}).
		Where("email", email).
		With("SocialAccounts").
		FirstOrFail(&user)

	if err != nil {
		if errors.Is(err, errors.OrmRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (s *UserService) HasSocialAccounts(user *models.User) bool {
	return len(user.SocialAccounts) > 0
}

func (s *UserService) GetSocialProviders(user *models.User) []string {
	var providers []string
	for _, account := range user.SocialAccounts {
		providers = append(providers, account.ProviderName)
	}
	return providers
}

func (s *UserService) GetUserWithActiveRoles(email string) (*models.User, error) {
	var user models.User
	err := facades.Orm().
		Query().
		Model(&models.User{}).
		Where("email", email).
		With("SocialAccounts").
		With("Roles", "is_active = ?", true).
		FirstOrFail(&user)

	if err != nil {
		if errors.Is(err, errors.OrmRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (s *UserService) HasActiveRole(user *models.User) bool {
	for _, role := range user.Roles {
		if role.IsActive {
			return true
		}
	}
	return false
}

func (s *UserService) CreateUserFromOAuthTx(tx orm.Query, data OAuthData) (*models.User, error) {
	user := &models.User{
		Name:      data.Name,
		Surname:   data.LastName,
		Email:     data.Email,
		AvatarURL: data.AvatarURL,
	}

	if err := tx.Create(&user); err != nil {
		return nil, err
	}

	if err := s.roleService.AssignRoleToUserTx(tx, user, models.RoleClient); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) GetUserByEmailTx(tx orm.Query, email string) (*models.User, error) {
	var user models.User

	err := tx.Model(&models.User{}).
		Where("email = ?", email).
		FirstOrFail(&user)

	if err != nil {
		if errors.Is(err, errors.OrmRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (s *UserService) HandleOrphanedSocialAccountTx(tx orm.Query, data OAuthData, orphanedUserID uuid.UUID) (*models.User, error) {
	var user models.User
	err := tx.WithTrashed().Where("id", orphanedUserID).First(&user)
	if err == nil {
		if user.DeletedAt.Valid {
			_, err := tx.Model(&user).Restore()
			if err != nil {
				return nil, fmt.Errorf("failed to restore soft-deleted user: %w", err)
			}

		}

		user.Name = data.Name
		user.Surname = data.LastName
		user.AvatarURL = data.AvatarURL
		if err := tx.Save(&user); err != nil {
			return nil, fmt.Errorf("failed to update restored user details: %w", err)
		}

		return &user, nil
	}

	if errors.Is(err, errors.OrmRecordNotFound) {

		newUser := &models.User{
			Name:      data.Name,
			Surname:   data.LastName,
			Email:     data.Email,
			AvatarURL: data.AvatarURL,
		}

		newUser.ID = orphanedUserID

		if err := tx.Create(&newUser); err != nil {
			return nil, fmt.Errorf("failed to recreate user with original ID: %w", err)
		}

		if err := s.roleService.AssignRoleToUserTx(tx, newUser, models.RoleClient); err != nil {
			return nil, fmt.Errorf("failed to assign role to recreated user: %w", err)
		}

		return newUser, nil
	}

	return nil, err
}

func (s *UserService) GetUserByID(id uuid.UUID) (*models.User, error) {
	var user models.User
	err := facades.Orm().
		Query().
		Model(&models.User{}).
		Where("id", id).
		FirstOrFail(&user)

	if err != nil {
		if errors.Is(err, errors.OrmRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}
