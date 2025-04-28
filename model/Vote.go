package model

type Vote struct {
	GormModel

	BlogId int  `json:"blog_id"`
	Blog   Blog `gorm:"foreignKey:BlogId"`

	AuthorId int  `json:"author_id"`
	Author   User `gorm:"foreignKey:AuthorId"`

	Value int `json:"value"`
}
