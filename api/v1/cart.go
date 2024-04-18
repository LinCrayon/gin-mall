package v1

import (
	"gin-mall/pkg/util/jwt"
	"gin-mall/pkg/util/log"
	"gin-mall/service"
	"gin-mall/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateCart(c *gin.Context) {
	claim, _ := jwt.ParseToken(c.GetHeader("Authorization"))
	var req types.CartCreateReq
	if err := c.ShouldBind(&req); err == nil {
		l := service.GetCartSrv()
		res := l.CreateCart(c.Request.Context(), claim.Id, &req)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		log.LogrusObj.Infoln(err)
	}
}
func ListCart(c *gin.Context) {
	var req types.CartListReq
	claim, _ := jwt.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&req); err == nil {
		l := service.GetCartSrv()
		res := l.ListCart(c.Request.Context(), claim.Id)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		log.LogrusObj.Infoln(err)
	}
}

func UpdateCart(c *gin.Context) {
	var req types.UpdateCartServiceReq
	if err := c.ShouldBind(&req); err == nil {
		l := service.GetCartSrv()
		res := l.UpdateCart(c.Request.Context(), c.Param("id"), &req)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		log.LogrusObj.Infoln(err)
	}
}

func DeleteCart(c *gin.Context) {
	var req types.CartDeleteReq
	if err := c.ShouldBind(&req); err == nil {
		l := service.GetCartSrv()
		res := l.DeleteCart(c.Request.Context(), &req)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		log.LogrusObj.Infoln(err)
	}
}
