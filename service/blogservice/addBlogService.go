package service

import (
	"blog/model"
	"blog/repository"
	"errors"
	"strconv"

	"gorm.io/gorm"
)

func AddBlogService(db *gorm.DB, blogName string, content string, category string, authorId string) (string, error) {
	if blogName == "" || content == "" {
		return "", errors.New("category name and description must not be empty")
	}

	iId, err := strconv.Atoi(category)
	if err != nil {
		return "", errors.New("invalid id value")
	}

	iAuthorId, err := strconv.Atoi(authorId)
	if err != nil{
		return"", errors.New("invalid Author id value")
	}

	blog := model.Blog{
		BlogTitle:   blogName,
		BlogContent: content,
		CategoryId: iId,
		AuthorId: iAuthorId,
		Status: "pending",
	}

	_,err = repository.AddBlog(model.DB, blog)
	if err != nil{
		return"", err
	}

	return "add Blog complete: ", nil
}
