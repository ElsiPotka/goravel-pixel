package models

import "github.com/google/uuid"

type SocialAccount struct {
	BaseModel
	UserID       uuid.UUID `gorm:"type:uuid;not null"`
	ProviderName string    `gorm:"type:varchar(50);not null"`
	ProviderID   string    `gorm:"type:varchar(255);not null"`

	User User `gorm:"foreignKey:UserID"`
}
