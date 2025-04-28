package model


type Category struct {
	GormModel
	CategoryName string `json:"categoryName"`
	Description  string `json:"description"`
}
