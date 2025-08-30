package models

type Status struct {
	BaseModel
	Name        string `json:"name" gorm:"type:varchar(255);not null"`
	Description string `json:"description" gorm:"type:text"`
	Color       string `json:"color" gorm:"type:varchar(255);not null"`
	Context     string `json:"context" gorm:"type:varchar(255);not null"`
}
