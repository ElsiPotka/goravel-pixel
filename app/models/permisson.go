package models

type Permission struct {
	BaseModel
	Permission  string `gorm:"type:varchar(50);unique;not null" json:"permission"`
	Description string `gorm:"type:text" json:"description"`
	IsActive    bool   `gorm:"default:true" json:"is_active"`

	Roles []Role `gorm:"many2many:role_has_permission;" json:"roles"`
}
