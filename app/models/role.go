package models

type Role struct {
	BaseModel
	Role        RoleType `gorm:"type:varchar(50);unique;not null" json:"role"`
	Description string   `gorm:"type:text" json:"description"`
	IsActive    bool     `gorm:"default:true" json:"is_active"`

	Users       []User       `gorm:"many2many:user_has_role;" json:"users"`
	Permissions []Permission `gorm:"many2many:role_has_permission;" json:"permissions"`
}
