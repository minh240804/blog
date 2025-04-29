package service

import (
	"blog/model"
	"blog/repository"
	"errors"
	"strconv"

	"gorm.io/gorm"
)

func Decentralization(db *gorm.DB, sId string) (string, error) {
	iId, err := strconv.Atoi(sId)
	if err != nil {
		return "", errors.New("invalid if value")
	}

	decentUser, err := repository.GetUserbyId(model.DB, uint(iId))
	if err != nil {
		return "", err
	}

	if decentUser.Role == "user"{
		decentUser.Role = "censor"
	} else{
		decentUser.Role = "user"
	}

	uid, err := repository.UpdateUser(model.DB, decentUser)
	if err != nil {
		return "", err
	}
	return "delete " + strconv.Itoa(int(uid)) + " user compelted", err
}
