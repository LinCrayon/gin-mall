package v1

import (
	"errors"
	"gin-mall/consts"
	"gin-mall/pkg/util/ctl"
	"gin-mall/pkg/util/jwt"
	"gin-mall/pkg/util/log"
	"gin-mall/service"
	"gin-mall/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserRegister() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.UserRegisterReq
		//请求体中的JSON数据解析并绑定到 req 结构体上
		if err := ctx.ShouldBind(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
			return
		}
		// 参数校验
		if req.Key == "" || len(req.Key) != consts.EncryptMoneyKeyLength {
			err := errors.New("key长度错误,必须是6位数")
			ctx.JSON(http.StatusOK, ErrorResponse(err))
			return
		}
		l := service.GetUserSrv()
		res := l.UserRegister(ctx.Request.Context(), &req)
		ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, res))
	}
}

func UserLogin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.UserLoginReq
		if err := ctx.ShouldBind(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}
		l := service.GetUserSrv()
		res := l.UserLogin(ctx.Request.Context(), &req)
		ctx.JSON(http.StatusOK, ctl.RespSuccess(ctx, res))
	}
}

func UserUpdate(c *gin.Context) {
	var userUpdate service.UserService
	//解析token
	claims, _ := jwt.ParseToken(c.GetHeader("Authorization")) //从请求头中提取的Authorization并进行解析
	if err := c.ShouldBind(&userUpdate); err == nil {
		res := userUpdate.UserUpdate(c.Request.Context(), claims.Id)
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
		log.LogrusObj.Infoln(err)
	}
}
func UploadAvatar(c *gin.Context) {
	// HTTP请求中获取上传的文件信息，包括文件内容和文件大小
	file, fileHeader, _ := c.Request.FormFile("file")
	fileSize := fileHeader.Size //文件的大小,字节为单位
	uploadAvatar := service.UserService{}
	claim, _ := jwt.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&uploadAvatar); err == nil {
		res := uploadAvatar.PostAvatar(c.Request.Context(), claim.Id, file, fileSize)
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
		log.LogrusObj.Infoln(err)
	}
}

// SendEmail 发送邮箱
func SendEmail(c *gin.Context) {
	var req types.SendEmailServiceReq
	claim, _ := jwt.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&req); err == nil {
		l := service.GetUserSrv()
		res := l.SendEmail(c.Request.Context(), claim.Id, &req)
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
		log.LogrusObj.Infoln(err)
	}
}

// ValidEmail  验证邮箱
func ValidEmail(c *gin.Context) {
	var req types.ValidEmailServiceReq
	if err := c.ShouldBind(&req); err == nil {
		l := service.GetUserSrv()
		res := l.ValidEmail(c.Request.Context(), c.GetHeader("Authorization"), &req)
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
		log.LogrusObj.Infoln(err)
	}
}
