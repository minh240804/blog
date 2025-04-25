package model

import (
	"time"
)

type User struct {
	UserId   int       `json:"user_id"`
	UserName string    `json:"user_name"`
	Password string    `json:"password"`
	Role     string    `json:"role"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
}
