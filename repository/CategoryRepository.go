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

func DeleteCategory(db *gorm.DB, category model.Category) (uint, error) {
	if err := db.Delete(&category).Error; err != nil {
		return 0, err
	}
	return category.ID, nil
}

func GetCategorybyId(db *gorm.DB, id uint) (model.Category, error) {
	var category model.Category
	if err := db.First(&category, id).Error; err != nil {
		return model.Category{}, err
	}
	return category, nil
}

func UpdateCategory(db *gorm.DB, category model.Category) (int64, error) {
	res := db.Save(&category)
	return res.RowsAffected, res.Error
}

func CountTotalCategory(db *gorm.DB, name string) (int64, error) {
	var count int64
	res := db.Model(&model.Category{}).Where("category_name like ?", name).Count(&count)
	return count, res.Error
}

func GetAllCategories(db *gorm.DB, page int, limit int, name string) ([]model.Category, error) {
	var sliceCategories = []model.Category{}
	res := db.Model(&model.Category{}).Where("category_name like ?", name).Limit(limit).Offset(limit * (page - 1)).Find(&sliceCategories)
	return sliceCategories, res.Error
}
