package model

type Blog struct{
	GormModel
	BlogTitle   string 	`json:"blogTitle"`
	BlogContent string 	`json:"blogContent"`

	CategoryId int		`json:categoryId`	
	Category   Category `gorm:"foreignKey:CategoryId"`

	AuthorId int		`json:authorId`
	Author   User `gorm:"foreignKey:AuthorId"`
}

