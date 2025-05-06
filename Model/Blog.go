package model

import (
	"time"
)

type Blog struct {
	GormModel
	BlogTitle   string `json:"blog_title"`
	BlogContent string `json:"blog_content"`

	CategoryId int      `json:"category_id"`
	Category   Category `gorm:"foreignKey:CategoryId"`

	AuthorId int  `json:"author_id"`
	Author   User `gorm:"foreignKey:AuthorId"`

	Comments []Comment `json:"comments"`

	Vote 	 []Vote `json:"vote"`

	Status string `gorm:"default:pending"`
}

type ApiBlog struct {
	BlogTitle   string    `json:"blog_title"`
	Author		User	`json:"author"`
	UpdatedAt  time.Time `json:"update_time"`
	Vote        float32   `json:"vote"`
	CommentCount	int 		`json:"comment_count"`
}
