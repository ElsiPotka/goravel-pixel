package models

import "github.com/google/uuid"

type Event struct {
	BaseModel
	Name        string     `json:"name" gorm:"type:varchar(255);not null"`
	Description string     `json:"description" gorm:"type:text"`
	TypeID      *uuid.UUID `json:"type_id"`
	Type        *Type      `json:"type" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
