package v1

import (
	"gin-mall/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ListCarousels  列表轮播
func ListCarousels(c *gin.Context) {
	var service service.CarouselsService
	//claim, _ := jwt.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&service); err == nil {
		//l := service.GetCarouselSrv()
		res := service.ListCarousel(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}
