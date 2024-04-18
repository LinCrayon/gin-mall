package service

import (
	"context"
	"fmt"
	"gin-mall/pkg/e"
	"gin-mall/pkg/util/log"
	"gin-mall/pkg/util/upload"
	"gin-mall/repository/db/dao"
	"gin-mall/repository/db/model"
	"gin-mall/serializer"
	"gin-mall/types"
	"mime/multipart"
	"strconv"
	"sync"
)

var ProductSrvIns *ProductSrv
var ProductSrvOnce sync.Once

type ProductSrv struct {
}

func GetProductSrv() *ProductSrv {
	ProductSrvOnce.Do(func() {
		ProductSrvIns = &ProductSrv{}
	})
	return ProductSrvIns
}
func (p *ProductSrv) ProductCreate(ctx context.Context, uid uint, files []*multipart.FileHeader, req *types.ProductCreateReq) serializer.Response {
	var boss *model.User
	code := e.SUCCESS
	boss, _ = dao.NewUserDao(ctx).GetUserById(uid)
	//第一张作为封面图
	tmp, err := files[0].Open()
	if err != nil {
		fmt.Println("第一张作为封面图err")
	}

	path, err := upload.ProductUploadToLocalStatic(tmp, uid, req.Name)
	if err != nil {
		code = e.ErrorProductImgUpload
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	product := &model.Product{
		Name:          req.Name,
		CategoryId:    req.CategoryId,
		Title:         req.Title,
		Info:          req.Info,
		ImgPath:       path,
		Price:         req.Price,
		DiscountPrice: req.DiscountPrice,
		Num:           req.Num,
		OnSale:        true,
		BossId:        uid,
		BossName:      boss.UserName,
		BossAvatar:    boss.Avatar,
	}
	productDao := dao.NewProductDao(ctx)
	err = productDao.CreateProduct(product)
	if err != nil {
		code = e.ERROR
		log.LogrusObj.Error(err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	wg := new(sync.WaitGroup)
	wg.Add(len(files))
	for index, file := range files {
		num := strconv.Itoa(index)
		productImgDao := dao.NewProductImgDaoByDB(productDao.DB) //DB复用
		tmp, err := file.Open()
		if err != nil {
			fmt.Println("open err")
		}
		path, err = upload.ProductUploadToLocalStatic(tmp, uid, req.Name+num)
		if err != nil {
			code = e.ErrorProductImgUpload
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
		productImg := model.ProductImg{
			ProductId: product.ID,
			ImgPath:   path,
		}
		err = productImgDao.CreateProductImg(&productImg)
		if err != nil {
			code = e.ERROR
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
		wg.Done()
	}
	wg.Wait()
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildProduct(product),
	}
}

// ListProduct 获取商品列表
func (p *ProductSrv) ListProduct(ctx context.Context, req *types.ProductListReq) serializer.Response {
	var products []*model.Product
	var err error
	code := e.SUCCESS
	if req.PageSize == 0 {
		req.PageSize = 15
	}
	condition := make(map[string]interface{})
	if req.CategoryId != 0 {
		condition["Category_id"] = req.CategoryId
	}
	productDao := dao.NewProductDao(ctx)
	total, err := productDao.CountProductByCondition(condition)
	if err != nil {
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		productDao = dao.NewProductDaoByDB(productDao.DB)
		products, _ = productDao.ListProductByCondition(condition, req.BasePage)
		wg.Done() //标记一个任务已完成
	}()
	wg.Wait() //阻塞当前 goroutine，直到 WaitGroup 中所有的等待任务都执行完成

	return serializer.BuildListResponse(serializer.BuildProducts(products), uint(total))
}

// SearchProduct 搜索商品
func (this *ProductSrv) SearchProduct(ctx context.Context, req *types.ProductSearchReq) serializer.Response {
	var err error
	code := e.SUCCESS
	if req.PageSize == 0 {
		req.PageSize = 15
	}
	productDao := dao.NewProductDao(ctx)
	products, count, err := productDao.SearchProduct(req.Info, model.BasePage(req.BasePage))
	if err != nil {
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.BuildListResponse(serializer.BuildProducts(products), uint(count))

}

// ShowProduct 查看商品详细详细
func (p *ProductSrv) ShowProduct(ctx context.Context, id string) serializer.Response {
	var err error
	code := e.SUCCESS
	pId, _ := strconv.Atoi(id)
	productDao := dao.NewProductDao(ctx)
	product, err := productDao.GetProductById(uint(pId))
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
		Data:   serializer.BuildProduct(product),
	}
}
