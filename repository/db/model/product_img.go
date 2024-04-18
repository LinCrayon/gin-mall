package model

import "github.com/jinzhu/gorm"

type ProductImg struct {
	gorm.Model
	ProductId uint `gorm:"not null"`
	ImgPath   string
}
