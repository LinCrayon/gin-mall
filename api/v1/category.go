package v1

import (
	"gin-mall/pkg/util/log"
	"gin-mall/service"
	"gin-mall/types"
	"github.com/gin-gonic/gin"
)

func ListCategories(c *gin.Context) {
	var req types.ListCategoryReq
	if err := c.ShouldBind(&req); err == nil {
		l := service.GetCategorySrv()
		res := l.List(c.Request.Context())
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
		log.LogrusObj.Infoln(err)
	}
}
