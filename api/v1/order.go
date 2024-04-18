package v1

import (
	"gin-mall/pkg/util/jwt"
	"gin-mall/pkg/util/log"
	"gin-mall/service"
	"gin-mall/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateOrder(c *gin.Context) {
	claim, _ := jwt.ParseToken(c.GetHeader("Authorization"))
	var req types.OrderCreateReq
	if err := c.ShouldBind(&req); err == nil {
		l := service.GetOrderSrv()
		res := l.CreateOrder(c.Request.Context(), claim.Id, &req)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		log.LogrusObj.Infoln(err)
	}
}
func ListOrders(c *gin.Context) {
	claim, _ := jwt.ParseToken(c.GetHeader("Authorization"))
	var req types.OrderListReq
	if err := c.ShouldBind(&req); err == nil {
		l := service.GetOrderSrv()
		res := l.ListOrder(c.Request.Context(), claim.Id, &req)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		log.LogrusObj.Infoln(err)
	}
}
func ShowOrder(c *gin.Context) {
	var req types.OrderShowReq
	if err := c.ShouldBind(&req); err == nil {
		l := service.GetOrderSrv()
		res := l.ShowOrder(c.Request.Context(), c.Param("id"), &req)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		log.LogrusObj.Infoln(err)
	}
}
func DeleteOrder(c *gin.Context) {
	var req types.OrderDeleteReq
	if err := c.ShouldBind(&req); err == nil {
		l := service.GetOrderSrv()
		res := l.DeleteOrder(c.Request.Context(), c.Param("id"), &req)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		log.LogrusObj.Infoln(err)
	}
}
