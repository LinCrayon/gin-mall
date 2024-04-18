package service

import (
	"context"
	"fmt"
	"gin-mall/repository/db/dao"
	"gin-mall/serializer"
	"strconv"
	"sync"
)

var ProductImgSrvIns *ProductImgSrv
var ProductImgSrvOnce sync.Once

type ProductImgSrv struct {
}

func GetProductImgSrv() *ProductImgSrv {
	ProductImgSrvOnce.Do(func() {
		ProductImgSrvIns = &ProductImgSrv{}
	})
	return ProductImgSrvIns
}

func (p *ProductImgSrv) ListProductImg(ctx context.Context, pId string) serializer.Response {
	productImgDao := dao.NewProductImgDao(ctx)
	productId, err := strconv.Atoi(pId)
	if err != nil {
		fmt.Println("转换失败:", err)
	} else {
		fmt.Println("转换成功:", productId)
	}
	productImgs, _ := productImgDao.ListProductImgByProductId(uint(productId))
	return serializer.BuildListResponse(serializer.BuildProductImgs(productImgs), uint(len(productImgs)))
}
