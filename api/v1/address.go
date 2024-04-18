package v1

import (
	"gin-mall/pkg/util/jwt"
	"gin-mall/pkg/util/log"
	"gin-mall/service"
	"gin-mall/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateAddress(c *gin.Context) {
	claim, _ := jwt.ParseToken(c.GetHeader("Authorization"))
	var req types.AddressCreateReq
	if err := c.ShouldBind(&req); err == nil {
		l := service.GetAddressSrv()
		res := l.CreateAddress(c.Request.Context(), claim.Id, &req)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		log.LogrusObj.Infoln(err)
	}
}
func GetAddress(c *gin.Context) {
	var req types.AddressGetReq
	if err := c.ShouldBind(&req); err == nil {
		l := service.GetAddressSrv()
		res := l.GetAddress(c.Request.Context(), c.Param("id"))
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		log.LogrusObj.Infoln(err)
	}
}

func ListAddress(c *gin.Context) {
	claim, _ := jwt.ParseToken(c.GetHeader("Authorization"))
	var req types.AddressListReq
	if err := c.ShouldBind(&req); err == nil {
		l := service.GetAddressSrv()
		res := l.ListAddress(c.Request.Context(), claim.Id, &req)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		log.LogrusObj.Infoln(err)
	}
}

func UpdateAddress(c *gin.Context) {
	claim, _ := jwt.ParseToken(c.GetHeader("Authorization"))
	var req types.AddressServiceReq
	if err := c.ShouldBind(&req); err == nil {
		l := service.GetAddressSrv()
		res := l.UpdateAddress(c.Request.Context(), claim.Id, c.Param("id"), &req)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		log.LogrusObj.Infoln(err)
	}
}

func DeleteAddress(c *gin.Context) {
	claim, _ := jwt.ParseToken(c.GetHeader("Authorization"))
	var req types.AddressDeleteReq
	if err := c.ShouldBind(&req); err == nil {
		l := service.GetAddressSrv()
		res := l.DeleteAddress(c.Request.Context(), claim.Id, c.Param("id"), &req)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		log.LogrusObj.Infoln(err)
	}
}
