package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/goravel/framework/database/orm"
	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	orm.SoftDeletes
}

func (base *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	if base.ID == uuid.Nil {
		u, err := uuid.NewV7()
		if err != nil {
			return err
		}
		base.ID = u
	}
	return nil
}
