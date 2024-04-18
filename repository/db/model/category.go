package model

import "gorm.io/gorm"

// todo 商品分类

type Category struct {
	gorm.Model
	CategoryName string
}
