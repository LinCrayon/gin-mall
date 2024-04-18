package serializer

import (
	"gin-mall/repository/db/model"
	"gin-mall/types"
)

type Address struct {
	ID       uint   `json:"id"`
	UserID   uint   `json:"user_id"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	Seen     bool   `json:"seen"`
	CreateAt int64  `json:"create_at"`
}

// BuildAddress 收货地址购物车
func BuildAddress(item *model.Address) types.AddressResp {
	return types.AddressResp{
		Id:      item.ID,
		UserId:  item.UserId,
		Name:    item.Name,
		Phone:   item.Phone,
		Address: item.Address,
		//Seen:      false,
		CreatedAt: item.CreatedAt.Unix(),
	}
}

// BuildAddresses 收货地址列表
func BuildAddresses(items []*model.Address) (addresses []types.AddressResp) {
	for _, item := range items {
		address := BuildAddress(item)
		addresses = append(addresses, address)
	}
	return addresses
}
