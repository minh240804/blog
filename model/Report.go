package model

type Report struct {
	GormModel

	BlogId int  `json:blogId`
	Blog   Blog `gorm:"foreignKey:BlogId"`

	AuthorId int  `json:authorId`
	Author   User `gorm:"foreignKey:AuthorId"`

	Content string
}
