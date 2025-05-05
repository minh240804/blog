package service

import (
	"blog/model"
	"blog/repository"
	"errors"
	"strconv"

	"gorm.io/gorm"
)

func UpdateBlogService(db *gorm.DB, blogName string, content string, categoryId string, authorId string, blogId string) (string, error) {
	iId, err := strconv.Atoi(categoryId)
	if err != nil {
		return "", err
	}

	iAuthorId, err := strconv.Atoi(authorId) // temp change later
	if err != nil {
		return "", err
	}

	iBlogId, err := strconv.Atoi(blogId)
	if err != nil {
		return "", err
	}

	upddateBlog, err := repository.GetBlogbyId(model.DB, uint(iBlogId))
	if err != nil {
		return "", err
	}

	if !checkAuthor(upddateBlog, iAuthorId) {
		return "", errors.New("you are not authorize to change this post")
	}

	if blogName != "" {
		upddateBlog.BlogTitle = blogName
	}

	if content != "" {
		upddateBlog.BlogContent = content
	}

	if iId != 0 {
		upddateBlog.CategoryId = iId
	}

	upddateBlog.Status = "pending"

	res, err := repository.UpdateBlog(model.DB, upddateBlog)
	if err != nil {
		return "", err
	}

	return "update Blog complete: " + strconv.Itoa(int(res)) + ". " + blogName, err
}

func checkAuthor(blog model.Blog, id int) bool {
	return uint(id) == uint(blog.AuthorId)
}

