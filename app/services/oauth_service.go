package services

import (
	"fmt"
	"pixel/app/models"

	"github.com/goravel/framework/contracts/database/orm"
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
)

type OAuthService struct {
	socialAccountService *SocialAccountService
	userService          *UserService
}

func NewOAuthService() *OAuthService {
	return &OAuthService{
		socialAccountService: NewSocialAccountService(),
		userService:          NewUserService(),
	}
}

type OAuthData struct {
	Provider     string `json:"provider"`
	ProviderID   string `json:"provider_id"`
	Email        string `json:"email"`
	Name         string `json:"name"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	AvatarURL    string `json:"avatar_url"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (s *OAuthService) ProcessOAuthUser(ctx http.Context, data OAuthData) (*models.User, string, string, int64, error) {
	if err := s.validateOAuthData(data); err != nil {
		return nil, "", "", 0, fmt.Errorf("invalid OAuth data: %w", err)
	}
	var user *models.User
	var err error

	err = facades.Orm().Transaction(func(tx orm.Query) error {
		user, err = s.processOAuthUserInTransaction(tx, data)
		return err
	})

	if err != nil {
		return nil, "", "", 0, err
	}

	return s.authenticateUser(ctx, user)
}

func (s *OAuthService) processOAuthUserInTransaction(tx orm.Query, data OAuthData) (*models.User, error) {
	socialAccount, user, err := s.socialAccountService.GetSocialAccountWithUserTx(tx, data.Email, data.Provider)

	if err == nil && socialAccount != nil && user == nil {
		user, err = s.userService.HandleOrphanedSocialAccountTx(tx, data, socialAccount.UserID)
		if err != nil {
			return nil, fmt.Errorf("failed to create user for existing social account: %w", err)
		}

		return user, nil
	}

	if socialAccount == nil {
		user, err = s.userService.GetUserByEmailTx(tx, data.Email)

		if err == nil && user != nil {
			_, err = s.socialAccountService.CreateSocialAccountTx(tx, user.ID, data.Provider, data.ProviderID)
			if err != nil {
				return nil, fmt.Errorf("failed to create social account for existing user: %w", err)
			}
			return user, nil
		}
	}

	if err == nil && user != nil && socialAccount != nil {

		return user, nil
	}

	user, err = s.userService.CreateUserFromOAuthTx(tx, data)
	if err != nil {
		return nil, fmt.Errorf("failed to create new user: %w", err)
	}

	_, err = s.socialAccountService.CreateSocialAccountTx(tx, user.ID, data.Provider, data.ProviderID)
	if err != nil {
		return nil, fmt.Errorf("failed to create social account for new user: %w", err)
	}

	return user, nil
}

func (s *OAuthService) authenticateUser(ctx http.Context, user *models.User) (*models.User, string, string, int64, error) {
	if !s.userService.HasActiveRole(user) {
		return nil, "", "", 0, fmt.Errorf("your account has no active role in the system. Please contact support for assistance")
	}

	accessToken, err := facades.Auth(ctx).Login(user)
	if err != nil {
		return nil, "", "", 0, fmt.Errorf("failed to generate access token: %w", err)
	}

	refreshToken := "-"
	// TODO: Implement refresh token generation
	// refreshToken, err := facades.Auth(ctx).Refresh()
	// if err != nil {
	//     return nil, "", "", 0, fmt.Errorf("failed to generate refresh token: %w", err)
	// }

	expiresIn := int64(facades.Config().GetInt("jwt.ttl") * 60)

	return user, accessToken, refreshToken, expiresIn, nil
}

func (s *OAuthService) validateOAuthData(data OAuthData) error {
	if data.Provider == "" {
		return fmt.Errorf("provider is required")
	}
	if data.ProviderID == "" {
		return fmt.Errorf("provider ID is required")
	}
	if data.Email == "" {
		return fmt.Errorf("email is required")
	}
	return nil
}
