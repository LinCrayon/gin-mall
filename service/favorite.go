package service

import (
	"context"
	"fmt"
	"gin-mall/pkg/e"
	"gin-mall/repository/db/dao"
	"gin-mall/repository/db/model"
	"gin-mall/serializer"
	"gin-mall/types"
	logging "github.com/sirupsen/logrus"
	"sync"
)

var FavoriteSrvIns *FavoriteSrv
var FavoriteSrvOnce sync.Once

type FavoriteSrv struct {
}

func GetFavoriteSrv() *FavoriteSrv {
	FavoriteSrvOnce.Do(func() {
		FavoriteSrvIns = &FavoriteSrv{}
	})
	return FavoriteSrvIns
}

func (f *FavoriteSrv) CreateFavorites(ctx context.Context, uId uint, req *types.FavoriteCreateReq) serializer.Response {
	code := e.SUCCESS
	favoriteDao := dao.NewFavoriteDao(ctx)
	exist, _ := favoriteDao.FavoriteExistOrNot(req.ProductId, uId) //判断用户是否存在
	if exist {
		code = e.ErrorExistFavorite
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserById(uId) //根据id查用户
	if err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "查询用户错误",
		}
	}
	bossDao := dao.NewUserDaoByDB(userDao.DB)
	fmt.Println(bossDao)
	boss, err := bossDao.GetUserById(req.BossId)
	if err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "查询boss错误",
		}
	}

	productDao := dao.NewProductDao(ctx)
	product, err := productDao.GetProductById(req.ProductId) //根据id查商品
	if err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "查询商品错误",
		}
	}

	favorite := &model.Favorite{
		UserId:    uId,
		User:      *user,
		ProductId: req.ProductId,
		Product:   *product,
		BossId:    req.BossId,
		Boss:      *boss,
	}
	favoriteDao = dao.NewFavoriteDaoByDB(favoriteDao.DB)
	err = favoriteDao.CreateFavorite(favorite) //创建收藏夹
	if err != nil {
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

func (f *FavoriteSrv) ShowFavorites(ctx context.Context, uid uint, req *types.FavoritesServiceReq) serializer.Response {
	favoritesDao := dao.NewFavoriteDao(ctx)
	code := e.SUCCESS
	if req.PageSize == 0 {
		req.PageSize = 15
	}
	favorites, total, err := favoritesDao.ListFavoriteByUserId(uid, req.PageSize, req.PageNum)
	if err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.BuildListResponse(serializer.BuildFavorites(ctx, favorites), uint(total))

}

func (f *FavoriteSrv) DeleteFavorites(ctx context.Context, req *types.FavoriteDeleteReq) serializer.Response {
	code := e.SUCCESS

	favoriteDao := dao.NewFavoriteDao(ctx)
	err := favoriteDao.DeleteFavoriteById(req.Id)
	if err != nil {
		logging.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Data:   e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Status: code,
		Data:   e.GetMsg(code),
	}
}
