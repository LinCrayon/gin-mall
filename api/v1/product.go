package v1

import (
	"gin-mall/pkg/util/jwt"
	"gin-mall/pkg/util/log"
	"gin-mall/service"
	"gin-mall/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateProduct 创建商品
func CreateProduct(c *gin.Context) {
	form, _ := c.MultipartForm() //解析请求体中的多部分表单数据
	files := form.File["file"]   //获取"file"字段对应的文件对象列表
	claim, _ := jwt.ParseToken(c.GetHeader("Authorization"))
	var req types.ProductCreateReq
	if err := c.ShouldBind(&req); err == nil {
		l := service.GetProductSrv()
		res := l.ProductCreate(c.Request.Context(), claim.Id, files, &req)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		log.LogrusObj.Infoln(err)
	}
}

// ListProduct 获取商品列表
func ListProduct(c *gin.Context) {
	var req types.ProductListReq
	if err := c.ShouldBind(&req); err == nil {
		l := service.GetProductSrv()
		res := l.ListProduct(c.Request.Context(), &req)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		log.LogrusObj.Infoln(err)
	}
}

// SearchProduct 搜索商品
func SearchProduct(c *gin.Context) {
	var req types.ProductSearchReq
	//claim, _ := jwt.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&req); err == nil {
		l := service.GetProductSrv()
		res := l.SearchProduct(c.Request.Context(), &req)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		log.LogrusObj.Infoln(err)
	}
}

func ShowProduct(c *gin.Context) {
	var req types.ProductShowReq
	if err := c.ShouldBind(&req); err == nil {
		l := service.GetProductSrv()
		res := l.ShowProduct(c.Request.Context(), c.Param("id"))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		log.LogrusObj.Infoln(err)
	}
}
