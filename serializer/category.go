package serializer

import (
	"gin-mall/repository/db/model"
	"gin-mall/types"
)

//TODO　类别

type Category struct {
	ID           uint   `json:"id"`
	CategoryName string `json:"category_name"`
	CreateAt     int64  `json:"create_at"`
}

func BuildCategory(item *model.Category) types.ListCategoryResp {
	return types.ListCategoryResp{
		ID:           item.ID,
		CategoryName: item.CategoryName,
		CreatedAt:    item.CreatedAt.Unix(),
	}
}

func BuildCategories(items []*model.Category) (categories []types.ListCategoryResp) {
	for _, item := range items {
		category := BuildCategory(item)
		categories = append(categories, category)
	}
	return categories
}
