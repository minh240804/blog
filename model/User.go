package model

type User struct {
	GormModel
	UserName string `json:"userName"`
	Password string `json:"password"`
	Role     string `json:"role"`
}
