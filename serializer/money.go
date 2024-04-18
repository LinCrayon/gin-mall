package serializer

import (
	"gin-mall/pkg/util"
	"gin-mall/repository/db/model"
	"gin-mall/types"
)

func BuildMoney(item *model.User, key string) types.ShowMoneyResp {
	util.Encrypt.SetKey(key)
	return types.ShowMoneyResp{
		UserId:    item.ID,
		UserName:  item.UserName,
		UserMoney: util.Encrypt.AesDecoding(item.Money),
	}
}
