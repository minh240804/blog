package repository

import (
	"blog/model"

	"gorm.io/gorm"
)

func AddCategory(db *gorm.DB, category model.Category) (uint, error) {
	if err := db.Create(&category).Error; err != nil {
		return 0, err
	}
	return category.ID, nil
}
