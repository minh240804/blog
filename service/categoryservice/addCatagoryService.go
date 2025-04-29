package service

import (
	"blog/model"
	"errors"

	"gorm.io/gorm"
)

func AddCategoryService(db *gorm.DB, categoryName string, descrpition string) (string, error) {
	if categoryName == "" || descrpition == "" {
		return "", errors.New("category name and description must not be empty")
	}

	category := model.Category{
		CategoryName: categoryName,
		Description:  descrpition,
	}

	if err := db.Create(&category).Error; err != nil {
		return "", err
	}

	return "add category complete: " + categoryName, nil
}
