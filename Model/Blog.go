package model

import (
	"time"
)
type Blog struct{
	BlogId int `json:"blog_id"`
	Title string `json:"title"`
	Content string `json:"content"`
	UserId int `json:"user_id"`
	CategoryId int `json:"category_id"`
	Status string `json:"status"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
}

