package model

import (
	"time"

	"gorm.io/gorm"
)

type GormModel struct {
	ID        uint           `gorm:"primarykey, index" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
