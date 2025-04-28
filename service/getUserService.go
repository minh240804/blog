package service

import (
	"blog/model"
)

type viewuser struct {
	Total      int          `json:"total"`
	Limit      int          `json:"limit"`
	Curencpage int          `json:"curencpage"`
	Users      []model.User `json:"Users"`
}

// func GetUserService(db *sql.DB, userName string, sLimit string, sPage string) (viewuser, error){
// 	if sPage == "" {
// 		sPage = "1"
// 	}
// 	if sLimit == "" {
// 		sLimit = "9"
// 	}
// 	if userName == "" {
// 		userName = "%"
// 	} else {
// 		userName = "%" + userName + "%"
// 	}

// 	iLimit, err := strconv.Atoi(sLimit)
// 	if err != nil{
// 		return viewuser{}, errors.New("Invalid limit value")
// 	}
// 	iPage, err := strconv.Atoi(sPage)
// 	if err != nil{
// 		return viewuser{}, errors.New("Invalid page value")
// 	}

// 	slicesUser, count, err := repository.GetAllUsers(db, iPage, iLimit)

// 	viewUser := viewuser{
// 		Total: count,
// 		Limit: iLimit,
// 		Curencpage: iPage,
// 		Users: slicesUser,
// 	}
// 	return viewUser, nil
// }
