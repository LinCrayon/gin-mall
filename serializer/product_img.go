package serializer

import (
	"gin-mall/config"
	"gin-mall/repository/db/model"
	"gin-mall/types"
)

type ProductImg struct {
	ProductID uint   `json:"product_id" form:"product_id"`
	ImgPath   string `json:"img_path" form:"img_path"`
}

func BuildProductImg(item *model.ProductImg) types.ProductImgResp {
	pImg := types.ProductImgResp{
		ProductId: item.ProductId,
		ImgPath:   config.Host + config.HttpPort + config.ProductPath + item.ImgPath,
	}
	return pImg
}

func BuildProductImgs(items []*model.ProductImg) (productImgs []types.ProductImgResp) {
	for _, item := range items {
		product := BuildProductImg(item)
		productImgs = append(productImgs, product)
	}
	return productImgs
}
