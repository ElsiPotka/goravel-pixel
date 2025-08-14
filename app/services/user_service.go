package services

import (
	"pixel/app/models"

	"github.com/goravel/framework/errors"
	"github.com/goravel/framework/facades"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
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
