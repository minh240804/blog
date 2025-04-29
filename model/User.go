package model

import "time"

type User struct {
	GormModel
	UserName string `json:"user_name" gorm:"unique"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type APIUser struct {
	Id        int       `json:"id"`
	UserName  string    `json:"user_name"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}