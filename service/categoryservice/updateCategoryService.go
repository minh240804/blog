package service

import (
	"blog/repository"
	"errors"
	"strconv"

	"gorm.io/gorm"
)

func UpdateCategory(db *gorm.DB, sId string, categoryName string, descrpition string) (string, error) {
	iId, err := strconv.Atoi(sId)
	if err != nil {
		return "", errors.New("invalid id value")
	}

	updateCategory, err := repository.GetCategorybyId(db, uint(iId))
	if err != nil{
		return "", err
	}

	if categoryName != ""{
		updateCategory.CategoryName = categoryName
	}

	if descrpition != ""{
		updateCategory.Description = descrpition
	}

	count, err := repository.UpdateCategory(db, updateCategory)
	if err != nil {
		return "", errors.New("invalid limit value")
	}

	return "update category complete: " + strconv.Itoa(int(count)) + ". " + categoryName, err
}
