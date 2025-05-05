package service

import (
	model "blog/model"
	"blog/repository"
	"errors"
	"strconv"

	"gorm.io/gorm"
)

func DeleteBlogService(db *gorm.DB, authorId string, blogId string) (string, error) {
	iAuthorId, err := strconv.Atoi(authorId)
	if err != nil {
		return "", err
	}

	iBlogId, err := strconv.Atoi(blogId)
	if err != nil {
		return "", err
	}

	deleteBlog, err := repository.GetBlogbyId(model.DB, uint(iBlogId))
	if err != nil {
		return "", err
	}

	if !checkDeleteBlog(deleteBlog, uint(iAuthorId)){
		return "", errors.New("you are not authorize to delete this post")
	}

	_, err = repository.DeleteBlog(model.DB, deleteBlog)
	if err != nil {
		return "", err
	}

	return "delete blog completed", nil
}

func checkDeleteBlog (blog model.Blog, uid uint) bool{
	user, err := repository.GetUserbyId(model.DB, uid)
	if err != nil {
		return false
	}

	if user.Role == "admin" || user.Role == "censor"{
		return true
	}

	return checkAuthor(blog, int(uid))
}