package service

import (
	"blog/model"
	"blog/repository"
	"errors"
	"strconv"

	"gorm.io/gorm"
)

type viewuser struct {
	Total       int             `json:"total"`
	Limit       int             `json:"limit"`
	CurencePage int             `json:"curence page"`
	Users       []model.User 	`json:"Users"`
}

func GetUserService(db *gorm.DB, userName string, sLimit string, sPage string, role string) (viewuser, error) {
	if sPage == "" {
		sPage = "1"
	}
	if sLimit == "" {
		sLimit = "9"
	}
	if userName == "" {
		userName = "%"
	} else {
		userName = "%" + userName + "%"
	}
	if role == "" {
		role = "%"
	} else {
		role = "%" + role + "%"
	}

	iLimit, err := strconv.Atoi(sLimit)
	if err != nil {
		return viewuser{}, errors.New("invalid limit value")
	}
	iPage, err := strconv.Atoi(sPage)
	if err != nil {
		return viewuser{}, errors.New("invalid page value")
	}

	//log.Println("from service " + userName )

	total, err := repository.CountTotalUsers(db, userName)

	if err != nil {
		return viewuser{}, err
	}

	if ((iPage - 1) * iLimit) > int(total) {
		iPage = 1
	}

	slicesUser, err := repository.GetAllUsers(db, iPage, iLimit, userName, role)
	if err != nil {
		return viewuser{}, err
	}

	viewUser := viewuser{
		Total:       int(total),
		Limit:       iLimit,
		CurencePage: iPage,
		Users:       slicesUser,
	}
	return viewUser, nil
}
