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

func UpdateBlog(db *gorm.DB, blog model.Blog) (int64, error) {
	res := db.Save(blog)
	return res.RowsAffected, res.Error
}

func DeleteBlog(db *gorm.DB, blog model.Blog) (uint, error) {
	if err := db.Delete(&blog).Error; err != nil {
		return 0, err
	}
	return blog.ID, nil
}

func CountTotalBlog(db *gorm.DB, name string, category string) (int64, error) {
	var count int64
	res := db.Model(&model.Blog{}).Where("blog_title like ?", name). /*category*/ Count(&count)
	return count, res.Error
}

func GetAllBlog(db *gorm.DB, name string, category string, page int, limit int) ([]model.Blog, error) {
	// var slicesBlog = []model.ApiBlog{}
	// res := db.Preload("Author").Model(&model.Blog{}).Where("blog_title like ?", name). /*category*/ Limit(limit).Offset(limit * (page - 1)).Find(&slicesBlog)
	// return slicesBlog, res.Error

	var slicesBlog = []model.Blog{}
	err := db.
			Preload("Author").
			// Preload("Comments").
			// Preload("Vote").
			Model(&model.Blog{}).
			Where("blog_title like ?", name).
			Limit(limit).
			Offset(limit * (page - 1)).
			Find(&slicesBlog).Error
	return slicesBlog, err
}
