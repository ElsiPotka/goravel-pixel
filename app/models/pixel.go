package models

import "github.com/google/uuid"

type Pixel struct {
	BaseModel
	Name          string  `json:"name" gorm:"type:varchar(255);not null"`
	Description   string  `json:"description" gorm:"type:text"`
	WebsiteLogo   string  `json:"website_logo" gorm:"type:varchar(255)"`
	WebsiteUrl    string  `json:"website_url" gorm:"type:varchar(255)"`
	Price         float64 `json:"price" gorm:"type:float(10,2)"`
	Currency      string  `json:"currency" gorm:"type:varchar(10)"`
	AudienceProof string  `json:"audience_proof" gorm:"type:varchar(255)"`

	SourceID   *uuid.UUID `json:"source_id"`
	Source     *Source    `json:"source" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CategoryID *uuid.UUID `json:"category_id"`
	Category   *Category  `json:"category" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
