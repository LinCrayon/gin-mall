package service

import (
	"context"
	"gin-mall/pkg/e"
	"gin-mall/repository/db/dao"
	"gin-mall/repository/db/model"
	"gin-mall/serializer"
	"gin-mall/types"
	"strconv"
	"sync"
)

var AddressSrvIns *AddressSrv
var AddressSrvOnce sync.Once

type AddressSrv struct {
}

func GetAddressSrv() *AddressSrv {
	AddressSrvOnce.Do(func() {
		AddressSrvIns = &AddressSrv{}
	})
	return AddressSrvIns
}

func (a *AddressSrv) CreateAddress(ctx context.Context, uId uint, req *types.AddressCreateReq) serializer.Response {
	code := e.SUCCESS
	var err error
	var address *model.Address
	addressDao := dao.NewAddressDao(ctx)
	address = &model.Address{
		UserId:  uId,
		Name:    req.Name,
		Phone:   req.Phone,
		Address: req.Address,
	}
	err = addressDao.CreateAddress(address)
	if err != nil {
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

func (a *AddressSrv) GetAddress(ctx context.Context, aId string) serializer.Response {
	addressId, _ := strconv.Atoi(aId)
	code := e.SUCCESS
	var err error

	addressDao := dao.NewAddressDao(ctx)
	address, err := addressDao.GetAddressByAid(uint(addressId))
	if err != nil {
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildAddress(address),
	}

}

func (a *AddressSrv) ListAddress(ctx context.Context, uId uint, req *types.AddressListReq) serializer.Response {
	code := e.SUCCESS
	var err error
	addressDao := dao.NewAddressDao(ctx)
	addressList, err := addressDao.ListAddressByUid(uId)
	if err != nil {
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildAddresses(addressList),
	}
}

func (a *AddressSrv) UpdateAddress(ctx context.Context, uId uint, aId string, req *types.AddressServiceReq) serializer.Response {
	code := e.SUCCESS
	var err error
	var address *model.Address
	addressDao := dao.NewAddressDao(ctx)
	address = &model.Address{
		UserId:  uId,
		Name:    req.Name,
		Phone:   req.Phone,
		Address: req.Address,
	}
	addressId, err := strconv.Atoi(aId)
	err = addressDao.UpdateAddressByUserId(uint(addressId), address)
	if err != nil {
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

func (a *AddressSrv) DeleteAddress(ctx context.Context, uId uint, aId string, req *types.AddressDeleteReq) serializer.Response {
	code := e.SUCCESS
	var err error
	addressDao := dao.NewAddressDao(ctx)
	addressId, err := strconv.Atoi(aId)
	err = addressDao.DeleteAddressByAid(uint(addressId), uId)
	if err != nil {
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}
