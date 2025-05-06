package model


type Category struct {
	GormModel
	CategoryName string `json:"categoryName" gorm:"unique"`
	Description  string `json:"description"`
}

type ApiCategory struct{ // type for the selecte and choose at list blog
	ID 				uint 	`json:"id"`
	CategoryName	string 	`json:"categoryName"`
}