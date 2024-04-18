package v1

import (
	"gin-mall/pkg/util/jwt"
	"gin-mall/pkg/util/log"
	"gin-mall/service"
	"gin-mall/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateFavorites(c *gin.Context) {
	claim, _ := jwt.ParseToken(c.GetHeader("Authorization"))
	var req types.FavoriteCreateReq
	if err := c.ShouldBind(&req); err == nil {
		l := service.GetFavoriteSrv()
		res := l.CreateFavorites(c.Request.Context(), claim.Id, &req)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		log.LogrusObj.Infoln(err)
	}
}

// ShowFavorites 收藏夹详情接口
func ShowFavorites(c *gin.Context) {
	var req types.FavoritesServiceReq
	claim, _ := jwt.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&req); err == nil {
		l := service.GetFavoriteSrv()
		res := l.ShowFavorites(c.Request.Context(), claim.Id, &req)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		log.LogrusObj.Infoln(err)
	}
}

// DeleteFavorites 删除收藏夹
func DeleteFavorites(c *gin.Context) {
	var req types.FavoriteDeleteReq
	if err := c.ShouldBind(&req); err == nil {
		l := service.GetFavoriteSrv()
		res := l.DeleteFavorites(c.Request.Context(), &req)
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusBadRequest, ErrorResponse(err))
		log.LogrusObj.Infoln(err)
	}
}
