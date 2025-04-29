package repository

import (
	"blog/model"
	"errors"

	"gorm.io/gorm"
	// "github.com/pelletier/go-toml/query"
)

func GetAllUsers(db *gorm.DB, page int, limit int, userName string, role string) ([]model.APIUser, error) {
	// check the limitedd
	var slicesUsers = []model.APIUser{}
	res := db.Model(&model.User{}).Where("user_name like ?", userName).Where("role like ?", role).Limit(limit).Offset(limit * (page - 1)).Find(&slicesUsers)
	return slicesUsers, res.Error
}

func GetUserbyId(db *gorm.DB, id uint) (model.User, error) {
	user := model.User{}
	res := db.Where("id = ?", id).Find(&user)
	if user.ID == 0{
		return user, errors.New("user not found")
	}
	return user, res.Error
}

func CountTotalUsers(db *gorm.DB, userName string) (int64, error) {
	var count int64
	db.Model(&model.User{}).Where("user_name like ?", userName).Count(&count)
	return count, nil
}

func AddUser(db *gorm.DB, user model.User) (int64, error) {

	//log.Println("user: " + user.UserName + " password: " + user.Password + " role: " + user.Role)

	if user.Role == "" {
		user.Role = "user"
	}
	res := db.Create(&user)

	return res.RowsAffected, res.Error
}

func UpdateUser(db *gorm.DB, user model.User) (int64, error) {
	res := db.Save(&user)
	return res.RowsAffected, res.Error
}

func DeleteUser(db *gorm.DB, user model.User) (int64, error) {
	res := db.Delete(user)
	return res.RowsAffected, res.Error
}

func GetAdmin(db *gorm.DB) ([]model.User, error) {
	var slicesAdmin = []model.User{}
	res := db.Where(&model.User{Role: "admin"}).Find(&slicesAdmin)

	return slicesAdmin, res.Error
}
