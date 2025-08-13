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
