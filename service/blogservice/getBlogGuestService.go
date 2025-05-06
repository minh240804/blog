package service

import (
	"blog/model"
	"blog/repository"
	"errors"
	"strconv"
)

type viewBlogList struct {
	Total       int                 `json:"total"`
	Limit       int                 `json:"limit"`
	CurencePage int                 `json:"curence page"`
	Blogs       []model.Blog        `json:"blogs"`
	Category    []model.ApiCategory `json:"category"`
}

func GetGuestBlogServcie(sLimit string, sPage string, blogName string, category string) (viewBlogList, error) {
	if sPage == "" {
		sPage = "1"
	}
	if sLimit == "" {
		sLimit = "9"
	}
	if blogName == "" {
		blogName = "%"
	} else {
		blogName = "%" + blogName + "%"
	}
	if category == "" {
		blogName = "%"
	} else {
		blogName = "%" + blogName + "%"
	}

	iLimit, err := strconv.Atoi(sLimit)
	if err != nil {
		return viewBlogList{}, errors.New("invalid limit value")
	}
	iPage, err := strconv.Atoi(sPage)
	if err != nil {
		return viewBlogList{}, errors.New("invalid page value")
	}

	total, err := repository.CountTotalBlog(model.DB, blogName, category)

	if err != nil {
		return viewBlogList{}, err
	}

	if ((iPage - 1) * iLimit) > int(total) {
		iPage = 1
	}

	slicesBlogs, err := repository.GetAllBlog(model.DB, blogName, category, iPage, iLimit)
	if err != nil {
		return viewBlogList{}, err
	}

	// slicesApiBlog, err := convertoApi(slicesBlogs)
	// if err != nil {
	// 	return viewBlogList{}, err
	// }

	slicesCategories, err := repository.GetAllApiCategory(model.DB)
	if err != nil {
		return viewBlogList{}, err
	}

	viewBlogList := viewBlogList{
		Total:       int(total),
		Limit:       iLimit,
		CurencePage: iPage,
		Blogs:       slicesBlogs,
		Category:    slicesCategories,
	}

	return viewBlogList, nil
}

// func convertoApi(blog []model.Blog) ([]model.ApiBlog, error) {
// 	var apiBlog = []model.ApiBlog{}
// 	for _, v := range blog {
// 		apiBlogCast := model.ApiBlog{
// 			BlogTitle: v.BlogTitle,
// 			Author:    v.Author,
// 			UpdatedAt: v.Author.UpdatedAt,
// 			// Vote:         caculateVote(v.Vote),
// 			// CommentCount: len(v.Comments),
// 		}
// 		apiBlog = append(apiBlog, apiBlogCast)
// 	}
// 	return apiBlog, nil
// }

// func caculateVote(vote []model.Vote) float32 {
// 	if vote == nil {
// 		return 0.0
// 	}
// 	res := float32(0)
// 	count := 0
// 	for _, v := range vote {
// 		count++
// 		res += float32(v.Value)
// 	}
// 	res = res / float32(count)
// 	return res
// }
