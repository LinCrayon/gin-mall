package service

import (
	"context"
	"gin-mall/pkg/e"
	"gin-mall/pkg/util/log"
	"gin-mall/repository/db/dao"
	"gin-mall/serializer"
	"sync"
)

var CarouselSrvIns *CarouselSrv
var CarouselSrvOnce sync.Once

type CarouselSrv struct {
}

func GetCarouselSrv() *CarouselSrv {
	CarouselSrvOnce.Do(func() {
		CarouselSrvIns = &CarouselSrv{}
	})
	return CarouselSrvIns
}

type CarouselsService struct {
}

// ListCarousel 列表
func (c *CarouselsService) ListCarousel(ctx context.Context) serializer.Response {
	//var carousels []model.Carousel
	code := e.SUCCESS
	carouselDao := dao.NewCarouselDao(ctx)
	carousels, err := carouselDao.ListCarousel()
	if err != nil {
		log.LogrusObj.Info("err", err)
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return serializer.BuildListResponse(serializer.BuildCarousels(carousels), uint(len(carousels)))
}
