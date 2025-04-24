package model

import(
)
type Blog struct{
	BlogId int `json:"blog_id"`
	Title string `json:"title"`
	Content string `json:"content"`
	UserId int `json:"user_id"`
	Status string `json:"status"`
	
}