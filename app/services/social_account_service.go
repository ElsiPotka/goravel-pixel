package services

import (
	"pixel/app/models"
	"strings"

	"github.com/google/uuid"
	"github.com/goravel/framework/contracts/database/orm"
	"github.com/goravel/framework/errors"
	"github.com/goravel/framework/facades"
	"github.com/markbates/goth"
)

type SocialAccountService struct{}

func NewSocialAccountService() *SocialAccountService {
	return &SocialAccountService{}
}

func (s *SocialAccountService) GetSocialAccount(email string, providerName string) (*models.SocialAccount, error) {
	var socialAccount models.SocialAccount

	err := facades.Orm().Query().
		Model(&models.SocialAccount{}).
		Join("users", "social_accounts.user_id", "=", "users.id").
		Where("users.email", email).
		Where("social_accounts.provider_name", providerName).
		First(&socialAccount)

	if err != nil {
		return nil, err
	}

	return &socialAccount, nil
}

type SocialAccountWithUser struct {
	SocialAccount models.SocialAccount
	User          models.User
}

func (s *SocialAccountService) GetSocialAccountWithUser(email string, providerName string) (*models.SocialAccount, *models.User, error) {
	var socialAccount models.SocialAccount

	err := facades.Orm().Query().
		Model(&models.SocialAccount{}).
		Join("users", "social_accounts.user_id", "=", "users.id").
		Where("users.email", email).
		Where("social_accounts.provider_name", providerName).
		First(&socialAccount)

	if err != nil {
		return nil, nil, err
	}

	var user models.User
	err = facades.Orm().Query().
		Model(&models.User{}).
		Where("id", socialAccount.UserID).
		First(&user)

	if err != nil {
		return &socialAccount, nil, err
	}

	return &socialAccount, &user, nil
}

func (s *SocialAccountService) GetSocialAccounts(email string) ([]models.SocialAccount, error) {
	var socialAccounts []models.SocialAccount

	err := facades.Orm().Query().
		Model(&models.SocialAccount{}).
		Join("users", "social_accounts.user_id", "=", "users.id").
		Where("users.email", email).
		Get(&socialAccounts)

	if err != nil {
		return nil, err
	}

	return socialAccounts, nil
}

func (s *SocialAccountService) CreateSocialAccount(userID uuid.UUID, providerName, providerID string) (*models.SocialAccount, error) {

	socialAccount := models.SocialAccount{
		UserID:       userID,
		ProviderName: providerName,
		ProviderID:   providerID,
	}

	if err := facades.Orm().Query().Create(&socialAccount); err != nil {
		return nil, err
	}

	return &socialAccount, nil
}

func (s *SocialAccountService) CreateSocialAccountTx(tx orm.Query, userID uuid.UUID, providerName, providerID string) (*models.SocialAccount, error) {

	socialAccount := models.SocialAccount{
		UserID:       userID,
		ProviderName: providerName,
		ProviderID:   providerID,
	}

	if err := tx.Create(&socialAccount); err != nil {
		return nil, err
	}

	return &socialAccount, nil
}

func (s *SocialAccountService) LinkSocialAccount(userID uuid.UUID, providerName, providerID string) error {
	var existing models.SocialAccount

	err := facades.Orm().Query().
		Model(&models.SocialAccount{}).
		Where("user_id", userID).
		Where("provider_name", providerName).
		First(&existing)

	if err == nil {
		return nil
	}

	_, createErr := s.CreateSocialAccount(userID, providerName, providerID)
	return createErr
}

func (s *SocialAccountService) CreateGothUserWithSocialAccount(gu goth.User) (*models.User, *models.SocialAccount, error) {
	firstName, lastName := gu.FirstName, gu.LastName
	if firstName == "" && lastName == "" {
		firstName, lastName = s.SplitName(gu.Name)
	}

	var user models.User
	err := facades.Orm().Query().
		Model(&models.User{}).
		Where("email", gu.Email).
		First(&user)
	if err != nil {
		user = models.User{
			Name:      firstName,
			Surname:   lastName,
			Email:     gu.Email,
			AvatarURL: gu.AvatarURL,
		}
		if createErr := facades.Orm().Query().Create(&user); createErr != nil {
			return nil, nil, createErr
		}
	}

	sa, err := s.LinkOrCreateSocialAccount(user.ID, gu.Provider, gu.UserID)
	if err != nil {
		return &user, nil, err
	}

	return &user, sa, nil
}

func (s *SocialAccountService) LinkOrCreateSocialAccount(userID uuid.UUID, providerName, providerID string) (*models.SocialAccount, error) {
	err := s.LinkSocialAccount(userID, providerName, providerID)
	if err == nil {
		var sa models.SocialAccount
		queryErr := facades.Orm().Query().
			Model(&models.SocialAccount{}).
			Where("user_id", userID).
			Where("provider_name", providerName).
			First(&sa)
		if queryErr != nil {
			return nil, queryErr
		}
		return &sa, nil
	}

	sa, createErr := s.CreateSocialAccount(userID, providerName, providerID)
	if createErr != nil {
		return nil, createErr
	}

	return sa, nil
}

func (s *SocialAccountService) SplitName(fullName string) (string, string) {
	parts := strings.Fields(fullName)
	if len(parts) == 0 {
		return "", ""
	}
	if len(parts) == 1 {
		return parts[0], ""
	}
	return parts[0], strings.Join(parts[1:], " ")
}

func (s *SocialAccountService) GetSocialAccountWithUserTx(tx orm.Query, email string, providerName string) (*models.SocialAccount, *models.User, error) {
	var socialAccount models.SocialAccount

	err := tx.Model(&models.SocialAccount{}).
		Join("users", "social_accounts.user_id", "=", "users.id").
		Where("users.email", email).
		Where("social_accounts.provider_name", providerName).
		FirstOrFail(&socialAccount)

	if err != nil {
		if errors.Is(err, errors.OrmRecordNotFound) {
			return nil, nil, nil
		}
		return nil, nil, err
	}

	var user models.User
	err = tx.Model(&models.User{}).
		Where("id", socialAccount.UserID).
		FirstOrFail(&user)

	if err != nil {
		if errors.Is(err, errors.OrmRecordNotFound) {
			return &socialAccount, nil, nil
		}
		return &socialAccount, nil, err
	}

	return &socialAccount, &user, nil
}
