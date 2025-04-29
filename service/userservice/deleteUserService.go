package service

import (
	"blog/model"
	"blog/repository"
	"errors"
	"strconv"

	"gorm.io/gorm"
)

func DeleteUserService(db *gorm.DB, sId string) (string, error) {

	iId, err := strconv.Atoi(sId)
	if err != nil {
		return "", errors.New("invalid if value")
	}

	deleteUser, err := repository.GetUserbyId(model.DB, uint(iId))
	if err != nil {
		return "", err
	}

	//log.Print("find the deleted: " + strconv.Itoa(int(deleteUser.ID)))

	uid, err := repository.DeleteUser(model.DB, deleteUser)
	if err != nil {
		return "", err
	}
	return "delete " + strconv.Itoa(int(uid)) + " user compelted", err
}
