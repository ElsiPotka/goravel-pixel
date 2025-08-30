package models

type Source struct {
	BaseModel
	Name        string `json:"name" gorm:"type:varchar(255);not null"`
	Description string `json:"description" gorm:"type:text"`
	MaxAudience int    `json:"max_audience"`
	Types       []Type `json:"types" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
