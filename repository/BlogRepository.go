package repository

import (
	"blog/model"

	"gorm.io/gorm"
)

func AddBlog(db *gorm.DB, blog model.Blog) (int64, error) {
	res := db.Create(&blog)
	return res.RowsAffected, res.Error
}

func GetBlogbyId(db *gorm.DB, id uint) (model.Blog, error) {
	blog := model.Blog{}
	res := db.Where("id = ?", id).First(&blog)
	return blog, res.Error
}

func UpdateBlog(db *gorm.DB, blog model.Blog) (int64, error){
	res := db.Save(blog)
	return res.RowsAffected, res.Error
}

func DeleteBlog(db *gorm.DB, blog model.Blog) (uint, error){
	if err := db.Delete(&blog).Error; err != nil {
		return 0, err
	}
	return blog.ID, nil
}