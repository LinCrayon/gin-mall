package serializer

import (
	"context"
	"fmt"
	"gin-mall/config"
	"gin-mall/repository/db/dao"
	"gin-mall/repository/db/model"
	"gin-mall/types"
)

// 购物车
type Cart struct {
	ID            uint   `json:"id"`
	UserID        uint   `json:"user_id"`
	ProductID     uint   `json:"product_id"`
	CreateAt      int64  `json:"create_at"`
	Num           uint   `json:"num"`
	MaxNum        uint   `json:"max_num"`
	Check         bool   `json:"check"`
	Name          string `json:"name"`
	ImgPath       string `json:"img_path"`
	DiscountPrice string `json:"discount_price"`
	BossId        uint   `json:"boss_id"`
	BossName      string `json:"boss_name"`
	Desc          string `json:"desc"`
}

func BuildCart(cart *model.Cart, product *model.Product, boss *model.User) types.CartResp {
	c := types.CartResp{
		ID:            cart.ID,
		UserID:        cart.UserId,
		ProductID:     cart.ProductId,
		CreatedAt:     cart.CreatedAt.Unix(),
		Num:           cart.Num,
		MaxNum:        cart.MaxNum,
		Check:         cart.Check,
		Name:          product.Name,
		ImgPath:       config.Host + config.HttpPort + config.ProductPath + product.ImgPath,
		DiscountPrice: product.DiscountPrice,
		BossId:        boss.ID,
		BossName:      boss.UserName,
	}
	//if config.UploadModel == consts.UploadModelOss {
	//	c.ImgPath = product.ImgPath
	//}

	return c
}

func BuildCarts(items []*model.Cart) (carts []types.CartResp) {
	for _, item1 := range items {
		product, err := dao.NewProductDao(context.Background()).
			GetProductById(item1.ProductId)
		if err != nil {
			fmt.Println("GetProductById err")
			continue
		}
		boss, err := dao.NewUserDao(context.Background()).
			GetUserById(item1.BossId)
		if err != nil {
			fmt.Println("GetUserById err")
			continue
		}
		cart := BuildCart(item1, product, boss)
		carts = append(carts, cart)
	}
	return carts
}
