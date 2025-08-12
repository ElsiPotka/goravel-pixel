package models

type User struct {
	BaseModel
	Name     string
	Email    string
	Password string

	Roles []Role `gorm:"many2many:user_has_role;" json:"roles"`
}
