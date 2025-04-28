package service

import (
	"blog/model"
	"blog/repository"
	"errors"
	"strconv"

	"gorm.io/gorm"
)

func AddUserService(db *gorm.DB, userName string, password string, role string) (string, error) {

	if userName == "" || password == "" {
		return "", errors.New("user name and password must not be empty.")
	}

	inputUser := model.User{
		UserName: userName,
		Password: password,
		Role:     role,
	}

	uid, err := repository.AddUser(model.DB, inputUser)
	if err != nil {
		return "", err
	}
	return "add User complete: " + strconv.Itoa(int(uid)) + ". " + userName, err

}
