package v1

import (
	"gin-mall/pkg/util/jwt"
	"gin-mall/pkg/util/log"
	"gin-mall/service"
	"gin-mall/types"
	"github.com/gin-gonic/gin"
)

// ShowMoney 获取用户金额
func ShowMoney(c *gin.Context) {
	var req types.ShowMoneyReq
	claim, _ := jwt.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&req); err == nil {
		l := service.GetMoneySrv()
		res := l.ShowMoney(c.Request.Context(), claim.Id, &req)
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
		log.LogrusObj.Infoln(err)
	}
}
