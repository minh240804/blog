package model


type Category struct {
	GormModel
	CategoryName string `json:"categoryName" gorm:"unique"`
	Description  string `json:"description"`
}
