package services

import (
	"fmt"
	"pixel/app/models"
	"time"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
)

type AuthService struct {
	roleService *RoleService
	userService *UserService
}

func NewAuthService() *AuthService {
	return &AuthService{
		roleService: NewRoleService(),
		userService: NewUserService(),
	}
}

type RegisterData struct {
	Name     string
	Surname  string
	Email    string
	Password string
}

type LoginData struct {
	Email    string
	Password string
}

func (s *AuthService) Register(ctx http.Context, data RegisterData) (*models.User, string, string, int64, error) {

	exists, err := s.userService.IsEmailTaken(data.Email)
	if err != nil {
		return nil, "", "", 0, err
	}
	if exists {
		return nil, "", "", 0, fmt.Errorf("email is already taken")
	}

	hashedPassword, err := facades.Hash().Make(data.Password)
	if err != nil {
		return nil, "", "", 0, err
	}

	now := time.Now()
	user := models.User{
		Name:      data.Name,
		Surname:   data.Surname,
		Email:     data.Email,
		Password:  hashedPassword,
		LastLogin: &now,
	}

	if err := facades.Orm().Query().Create(&user); err != nil {
		return nil, "", "", 0, err
	}

	var role models.Role
	if err := facades.Orm().Query().Where("role = ?", models.RoleClient).First(&role); err == nil {
		facades.Orm().Query().Model(&user).Association("Roles").Append(&role)
	}

	accessToken, err := facades.Auth(ctx).Login(&user)
	if err != nil {
		return nil, "", "", 0, err
	}

	_, err = facades.Auth(ctx).Parse(accessToken)
	if err != nil {
		return nil, "", "", 0, err
	}

	// refreshToken, err := facades.Auth(ctx).Refresh()
	// if err != nil {
	// 	return nil, "", "", 0, err
	// }

	refreshToken := " - "

	expiresIn := int64(facades.Config().GetInt("jwt.ttl") * 60)

	return &user, accessToken, refreshToken, expiresIn, nil
}

func (s *AuthService) Login(ctx http.Context, data LoginData) (*models.User, string, string, int64, error) {
	user, err := s.userService.GetUserByEmail(data.Email)
	if err != nil {
		return nil, "", "", 0, err
	}
	if user == nil {
		return nil, "", "", 0, fmt.Errorf("user with provided email not found")
	}

	match := facades.Hash().Check(data.Password, user.Password)
	if !match {
		return nil, "", "", 0, fmt.Errorf("incorrect password")
	}

	accessToken, err := facades.Auth(ctx).Login(user)
	if err != nil {
		return nil, "", "", 0, err
	}

	refreshToken := "-"
	// refreshToken, err := facades.Auth(ctx).Refresh()
	// if err != nil {
	//     return nil, "", "", 0, err
	// }

	expiresIn := int64(facades.Config().GetInt("jwt.ttl") * 60)

	return user, accessToken, refreshToken, expiresIn, nil
}
