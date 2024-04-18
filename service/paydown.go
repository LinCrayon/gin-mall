package service

import (
	"context"
	"errors"
	"fmt"
	"gin-mall/pkg/e"
	"gin-mall/pkg/util"
	"gin-mall/repository/db/dao"
	"gin-mall/repository/db/model"
	"gin-mall/serializer"
	"gin-mall/types"
	"strconv"
	"sync"
)

var PayDownSrvIns *PayDownSrv
var PayDownSrvOnce sync.Once

type PayDownSrv struct {
}

func GetPayDownSrv() *PayDownSrv {
	PayDownSrvOnce.Do(func() {
		PayDownSrvIns = &PayDownSrv{}
	})
	return PayDownSrvIns
}
func (this *PayDownSrv) PayDown(ctx context.Context, uid uint, req *types.PaymentDownReq) serializer.Response {
	util.Encrypt.SetKey(req.Key)
	code := e.SUCCESS
	orderDao := dao.NewOrderDao(ctx)
	tx := orderDao.Begin()
	order, err := orderDao.GetOrderById(req.OrderId)
	if err != nil {
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	money := order.Money
	num := order.Num
	money = money * float64(num)

	//用户扣钱
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserById(uid)

	if err != nil {
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	//对钱进行解密，减订单，加密保存
	moneyStr := util.Encrypt.AesDecoding(user.Money)
	moneyFloat, _ := strconv.ParseFloat(moneyStr, 64)

	if moneyFloat-money < 0.0 {
		tx.Rollback()
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  errors.New("金额不足").Error(),
		}
	}
	finMoney := fmt.Sprintf("%f", moneyFloat-money)
	user.Money = util.Encrypt.AesEncoding(finMoney)

	userDao = dao.NewUserDaoByDB(userDao.DB)
	err = userDao.UpdateUserById(uid, user)

	if err != nil {
		tx.Rollback()
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	//商家加钱
	var boss *model.User

	boss, err = userDao.GetUserById(req.BossId)
	if err != nil {
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	moneyStr = util.Encrypt.AesDecoding(boss.Money)
	moneyFloat, _ = strconv.ParseFloat(moneyStr, 64)

	finMoney = fmt.Sprintf("%f", moneyFloat+money)
	boss.Money = util.Encrypt.AesEncoding(finMoney)

	err = userDao.UpdateUserById(req.BossId, boss)
	if err != nil { // 更新boss金额失败，回滚
		tx.Rollback()
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	//扣商品
	var product *model.Product
	productDao := dao.NewProductDaoByDB(tx)
	product, err = productDao.GetProductById(req.ProductId)
	product.Num -= num
	if err != nil {
		tx.Rollback()
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	product.Num -= num
	err = productDao.UpdateProduct(uint(req.ProductId), product)
	if err != nil { // 更新商品数量减少失败，回滚
		tx.Rollback()
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	// 更新订单状态
	err = orderDao.DeleteOrderById(req.OrderId)
	if err != nil {
		tx.Rollback()
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	productUser := model.Product{
		Name:          product.Name,
		CategoryId:    product.CategoryId,
		Title:         product.Title,
		Info:          product.Info,
		ImgPath:       product.ImgPath,
		Price:         product.Price,
		DiscountPrice: product.DiscountPrice,
		Num:           num,
		OnSale:        false,
		BossId:        uid,
		BossName:      user.UserName,
		BossAvatar:    user.Avatar,
	}
	err = productDao.CreateProduct(&productUser)
	if err != nil {
		tx.Rollback()
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	tx.Commit()
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}
