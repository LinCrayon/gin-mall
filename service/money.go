package service

import (
	"context"
	"gin-mall/pkg/e"
	"gin-mall/repository/db/dao"
	"gin-mall/serializer"
	"gin-mall/types"
	"sync"
)

var MoneySrvIns *MoneySrv
var MoneySrvOnce sync.Once

type MoneySrv struct {
}

func GetMoneySrv() *MoneySrv {
	MoneySrvOnce.Do(func() {
		MoneySrvIns = &MoneySrv{}
	})
	return MoneySrvIns
}

// ShowMoney  展示用户的金额
func (m *MoneySrv) ShowMoney(ctx context.Context, uid uint, req *types.ShowMoneyReq) serializer.Response {
	code := e.SUCCESS
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserById(uid)
	if err != nil {
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildMoney(user, req.Key),
		Msg:    e.GetMsg(code),
	}
}
