package model

type Blog struct{
	GormModel
	BlogTitle   string 	`json:"blogTitle"`
	BlogContent string 	`json:"blogContent"`

	CategoryId 	int		`json:"category_id"`	
	Category   	Category `gorm:"foreignKey:CategoryId"`

	AuthorId 	int		`json:"author_id"`
	Author   	User 	`gorm:"foreignKey:AuthorId"`

	Status 		string	`gorm:"default:pending"`
}

