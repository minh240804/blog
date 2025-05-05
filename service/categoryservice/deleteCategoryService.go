package service

import (
	"blog/model"
	"blog/repository"
	"errors"
	"strconv"

	"gorm.io/gorm"
)

func DeleteCategoryService(db *gorm.DB, sId string) (string, error) {
	iId, err := strconv.Atoi(sId)
	if err != nil {
		return "", errors.New("invalid if value")
	}

	deleteCategory, err := repository.GetCategorybyId(model.DB, uint(iId))
	if err != nil {
		return "", err
	}

	message, err := repository.DeleteCategory(model.DB, deleteCategory)
	if err != nil {
		return "", err
	}
	return "delete " + strconv.Itoa(int(message)) + " category completed", err
}
