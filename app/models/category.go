package models

type Category struct {
	BaseModel
	Name        string  `json:"name" gorm:"type:varchar(255);not null"`
	Description string  `json:"description" gorm:"type:text"`
	Pixels      []Pixel `json:"pixels" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
