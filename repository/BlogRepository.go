package repository

import (
	"blog/model"

	"gorm.io/gorm"
)

func AddBlog(db *gorm.DB, blog model.Blog) (int64, error) {
	res := db.Create(&blog)
	return res.RowsAffected, res.Error
}
