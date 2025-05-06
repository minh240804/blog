package model

import "time"

type User struct {
	GormModel `json:"-"`
	UserName string `json:"user_name" gorm:"unique"`
	Password string `json:"-"`
	Role     string `json:"role"`
}

type APIUser struct {
	Id        int       `json:"id"`
	UserName  string    `json:"user_name"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
