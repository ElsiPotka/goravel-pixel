package models

import (
	"time"
)

type User struct {
	BaseModel
	Name           string          `json:"name"`
	Surname        string          `json:"surname"`
	Email          string          `gorm:"unique" json:"email"`
	Password       string          `gorm:"nullable" json:"-"`
	AvatarURL      string          `gorm:"type:varchar(255);nullable" json:"avatar_url"`
	LastLogin      *time.Time      `gorm:"type:timestamptz;nullable" json:"last_login"`
	Roles          []Role          `gorm:"many2many:user_has_role;" json:"roles"`
	SocialAccounts []SocialAccount `gorm:"foreignKey:UserID" json:"social_accounts"`
}
