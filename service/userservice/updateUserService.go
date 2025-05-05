package service

import (
	"blog/model"
	"blog/repository"
	"errors"
	"strconv"

	"gorm.io/gorm"
)

func UpdateUserService(db *gorm.DB, userName string, password string, role string, sId string) (string, error) {

	if userName == "" || password == "" {
		return "", errors.New("user name and password must not be empty")
	}

	iId, err := strconv.Atoi(sId)
	if err != nil {
		return "", errors.New("invalid id value")
	}

	UpdateUser, err := repository.GetUserbyId(model.DB, uint(iId))
	if err != nil {
		return "", err
	}

	//log.Println("user found: " + strconv.Itoa(int(UpdateUser.ID)))

	if userName != "" {
		UpdateUser.UserName = userName
	}
	if password != "" {
		UpdateUser.Password = password
	}
	if role != "" {
		UpdateUser.Role = role
	}

	count, err := repository.UpdateUser(model.DB, UpdateUser)
	if err != nil {
		return "", errors.New("invalid limit value")
	}

	return "update User complete: " + strconv.Itoa(int(count)) + ". " + userName, err
}
