package model

type User struct {
	GormModel
	UserName string `json:"user_name"`
	Password string `json:"password"`
	Role     string `json:"role"`
}
