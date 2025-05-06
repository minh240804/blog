package model

type User struct {
	GormModel
	UserName string `json:"user_name" gorm:"unique"`
	Password string `json:"-"`
	Role     string `json:"role"`
}
