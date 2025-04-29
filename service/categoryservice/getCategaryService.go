package service

import (
	"blog/model"

	"gorm.io/gorm"
)

func GetCategoryService(db *gorm.DB, sLimit string, sPage string, name string) ([]model.Category, error){
	return []model.Category{}, nil
}
