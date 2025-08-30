package models

import "github.com/google/uuid"

type Type struct {
	BaseModel
	Name        string     `json:"name" gorm:"type:varchar(255);not null"`
	Description string     `json:"description" gorm:"type:text"`
	SourceID    *uuid.UUID `json:"source_id"`
	Source      *Source    `json:"source" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Events      []Event    `json:"events" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
