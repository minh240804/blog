package model

import (
	"time"
)

type Category struct {
	CategoryId   int       `json:"category_id"`
	CategoryName string    `json:"category_name"`
	Description  string    `json:"description"`
	CreateAt     time.Time `json:"create_at"`
	UpdateAt     time.Time `json:"update_at"`
}
