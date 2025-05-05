package service

import (
	"blog/model"
	"blog/repository"
	"errors"
	"strconv"

	"gorm.io/gorm"
)

type viewcategory struct {
	Total       int              `json:"total"`
	Limit       int              `json:"limit"`
	CurencePage int              `json:"curence page"`
	Category    []model.Category `json:"Categories"`
}

func GetCategoryService(db *gorm.DB, sLimit string, sPage string, name string) (viewcategory, error) {
	//return []model.Category{}, nil
	if sPage == "" {
		sPage = "1"
	}
	if sLimit == "" {
		sLimit = "9"
	}
	if name == "" {
		name = "%"
	} else {
		name = "%" + name + "%"
	}

	iLimit, err := strconv.Atoi(sLimit)
	if err != nil {
		return viewcategory{}, errors.New("invalid limit value")
	}
	iPage, err := strconv.Atoi(sPage)
	if err != nil {
		return viewcategory{}, errors.New("invalid page value")
	}


	total, err := repository.CountTotalCategory(model.DB, name)

	if err != nil {
		return viewcategory{}, err
	}

	if ((iPage - 1) * iLimit) > int(total) {
		iPage = 1
	}

	slicesCategory, err := repository.GetAllCategories(model.DB, iPage, iLimit, name)
	if err != nil{
		return viewcategory{}, err
	}

	viewcategory := viewcategory{
		Total:   	int(total),
		Limit:   	iLimit,
		CurencePage:iPage,
		Category: slicesCategory,
	}

	return viewcategory, nil
}
