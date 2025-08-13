package models

type User struct {
	BaseModel
	Name           string
	Email          string
	Password       string          `gorm:"nullable" json:"-"`
	AvatarURL      string          `gorm:"type:varchar(255);nullable" json:"avatar_url"`
	Roles          []Role          `gorm:"many2many:user_has_role;" json:"roles"`
	SocialAccounts []SocialAccount `gorm:"foreignKey:UserID" json:"social_accounts"`
}
