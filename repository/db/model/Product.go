package model

import (
	"gin-mall/repository/cache"
	"gorm.io/gorm"
	"strconv"
)

// Product 商品模型
type Product struct {
	gorm.Model
	Name          string `gorm:"size:255;index"`
	CategoryId    uint
	Title         string
	Info          string `gorm:"size:1000"`
	ImgPath       string
	Price         string
	DiscountPrice string //打折后价格
	OnSale        bool   `gorm:"default:false"` //在售
	Num           int
	BossId        uint
	BossName      string
	BossAvatar    string
}

// View 获取点击数
func (product *Product) View() uint64 {
	countStr, _ := cache.RedisClient.Get(cache.ProductViewKey(product.ID)).Result()
	count, _ := strconv.ParseUint(countStr, 10, 64)
	return count
}
