package services

import (
	"fmt"
	"pixel/app/models"
	"strings"
	"time"

	"github.com/google/uuid"
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

	_ = s.roleService.AssignRoleToUser(&user, models.RoleClient)

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
	user, err := s.userService.GetUserWithActiveRoles(data.Email)
	if err != nil {
		return nil, "", "", 0, err
	}

	if user == nil {
		return nil, "", "", 0, fmt.Errorf("user with provided email not found")
	}

	if !s.userService.HasActiveRole(user) {
		return nil, "", "", 0, fmt.Errorf("your account has been deactivated. Please contact support for assistance")
	}

	if user.Password == "" {
		if s.userService.HasSocialAccounts(user) {
			providers := s.userService.GetSocialProviders(user)
			providerList := strings.Join(providers, " or ")
			return nil, "", "", 0, fmt.Errorf("this account was created using %s. Please use %s to login", providerList, providerList)
		} else {
			return nil, "", "", 0, fmt.Errorf("account setup incomplete. Please contact support")
		}
	}

	match := facades.Hash().Check(data.Password, user.Password)
	if !match {
		if s.userService.HasSocialAccounts(user) {
			providers := s.userService.GetSocialProviders(user)
			providerList := strings.Join(providers, " or ")
			return nil, "", "", 0, fmt.Errorf("incorrect password. You can try again or login using %s", providerList)
		}
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

func (s *AuthService) Logout(ctx http.Context) error {
	return facades.Auth(ctx).Logout()
}

func (s *AuthService) Impersonate(ctx http.Context, userID uuid.UUID) (*models.User, string, string, int64, error) {
	user, err := s.userService.GetUserByID(userID)
	if err != nil {
		return nil, "", "", 0, err
	}

	if user == nil {
		return nil, "", "", 0, fmt.Errorf("user not found")
	}

	accessToken, err := facades.Auth(ctx).Login(user)
	if err != nil {
		return nil, "", "", 0, err
	}

	refreshToken := "-"
	expiresIn := int64(facades.Config().GetInt("jwt.ttl") * 60)

	return user, accessToken, refreshToken, expiresIn, nil
}
