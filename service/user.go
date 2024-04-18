package service

import (
	"context"
	"fmt"
	"gin-mall/config"
	"gin-mall/consts"
	"gin-mall/pkg/e"
	"gin-mall/pkg/util"
	"gin-mall/pkg/util/jwt"
	"gin-mall/pkg/util/upload"
	"gin-mall/repository/db/dao"
	"gin-mall/repository/db/model"
	"gin-mall/serializer"
	"gin-mall/types"
	"gopkg.in/mail.v2"
	"mime/multipart"
	"strings"
	"time"

	"sync"
)

type UserSrv struct{}     //空结构体，表示该服务
var UserSrvIns *UserSrv   //全局变量，保存了UserSrv类型的单一实例
var UserSrvOnce sync.Once //sync.Once类型的实例。它用于确保只有在多个goroutine同时调用GetUserSrv时，UserSrvIns变量的初始化仅执行一次。

func GetUserSrv() *UserSrv { //返回UserSrv类型的单例实例。它使用sync.Once对象确保仅执行一次Do函数内部的初始化代码。
	UserSrvOnce.Do(func() { //在Do函数内部，创建一个新的UserSrv实例并将其赋值给UserSrvIns变量。这个初始化操作保证仅执行一次，无论GetUserSrv被并发调用多少次。
		UserSrvIns = &UserSrv{}
	})
	return UserSrvIns
}

//整个生命周期内只有一个UserSrv类型的实例，并以线程安全的方式处理初始化。sync.Once确保初始化代码仅执行一次

// UserService 管理用户注册服务
type UserService struct {
	NickName string `form:"nick_name" json:"nick_name"`
	UserName string `form:"user_name" json:"user_name"`
	Password string `form:"password" json:"password"`
	Key      string `form:"key" json:"key"` // 前端进行判断
}

// UserRegister 用户注册
func (u *UserSrv) UserRegister(ctx context.Context, req *types.UserRegisterReq) serializer.Response {
	code := e.SUCCESS
	//密文存储 对对称加密 //设置密钥
	util.Encrypt.SetKey(req.Key)

	userDao := dao.NewUserDao(ctx)
	_, exist, err := userDao.ExistOrNotByUserName(req.UserName)
	if err != nil {
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	if exist {
		code = e.ErrorExistUser
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	user := &model.User{
		NickName: req.NickName,
		UserName: req.UserName,
		Status:   model.Active,
		Avatar:   "avatar.JPG",
		Money:    util.Encrypt.AesEncoding(consts.UserInitMoney), //初始金额的加密
	}
	//密码的加密
	if err = user.SetPassword(req.Password); err != nil {
		code = e.ErrorFailEncryption
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	// 创建用户
	err = userDao.CreateUser(user)
	if err != nil {
		code = e.ERROR
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

// UserLogin 用户登录
func (this *UserSrv) UserLogin(ctx context.Context, req *types.UserLoginReq) serializer.Response {
	code := e.SUCCESS
	userDao := dao.NewUserDao(ctx)
	//判断用户是否存在
	user, exist, err := userDao.ExistOrNotByUserName(req.UserName) //exist=true是存在该用户
	if !exist || err != nil {
		code = e.ErrorNotExistUser
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "用户不存在,请先注册",
		}
	}
	//校验密码
	if user.CheckPassword(req.Password) == false {
		code = e.ErrorNotCompare
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "用户名密码错误,请重输入",
		}
	}
	// http 无状态 token 签发 生成JWT令牌
	token, err := jwt.GenerateToken(user.ID, req.UserName, 0)
	if err != nil {
		code = e.ErrorAuthToken
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "token验证失败",
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code), //BuildUser()序列化
		Data:   serializer.TokenData{User: serializer.BuildUser(user), Token: token},
	}
}

// UserUpdate 更新昵称
func (service *UserService) UserUpdate(ctx context.Context, uid uint) serializer.Response {
	var user *model.User
	var err error
	code := e.SUCCESS
	//找到用户
	userDao := dao.NewUserDao(ctx)
	user, err = userDao.GetUserById(uid)
	if err != nil {
		fmt.Println("GetUserById查询user失败")
	}

	//修改nick_name
	if service.NickName != "" {
		user.NickName = service.NickName
	}

	err = userDao.UpdateUserById(uid, user) //根据uid更新信息
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
		Data:   serializer.BuildUser(user),
	}
}

// PostAvatar 上传头像
func (service *UserService) PostAvatar(ctx context.Context, uid uint, file multipart.File, fileSize int64) serializer.Response {
	var user *model.User
	code := e.SUCCESS
	var err error
	userDao := dao.NewUserDao(ctx)
	user, err = userDao.GetUserById(uid)
	if err != nil {
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	//保存图片到本地
	path, err := upload.AvatarUploadToLocalStatic(file, uid, user.UserName)
	if err != nil {
		code = e.ErrorUploadFile
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	user.Avatar = path
	err = userDao.UpdateUserById(uid, user)
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
		Data:   serializer.BuildUser(user),
	}
}

// SendEmail 发送邮箱
func (u *UserSrv) SendEmail(ctx context.Context, uid uint, req *types.SendEmailServiceReq) serializer.Response {
	code := e.SUCCESS
	var address string
	var notice *model.Notice
	token, err := jwt.GenerateEmailToken(uid, req.OperationType, req.Email, req.Password)
	if err != nil {
		code = e.ErrorAuthToken
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	noticeDao := dao.NewNoticeDao(ctx)
	notice, err = noticeDao.GetNoticeById(req.OperationType)
	if err != nil {
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	address = config.ValidEmail + token //发送方
	mailStr := notice.Text
	//字符串mailStr中的所有匹配项 "Email" 替换为指定的 address
	mailText := strings.Replace(mailStr, "Email", address, -1)
	m := mail.NewMessage()                //创建消息对象
	m.SetHeader("From", config.SmtpEmail) //发件人
	m.SetHeader("To", req.Email)          //收件人地址
	m.SetHeader("Subject", "lsq")         //邮件的主题
	m.SetBody("text/html", mailText)      //邮件正文内容
	d := mail.NewDialer(config.SmtpHost, 465, config.SmtpEmail, config.SmtpPass)
	d.StartTLSPolicy = mail.MandatoryStartTLS //启用 TLS 加密的策略
	if err := d.DialAndSend(m); err != nil {
		code = e.ErrorSendEmail
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

func (u *UserSrv) ValidEmail(ctx context.Context, token string, req *types.ValidEmailServiceReq) serializer.Response {
	var userId uint
	var email string
	var password string
	var operationType uint
	code := e.SUCCESS
	//验证token
	if token == "" {
		code = e.InvalidParams
	} else {
		claims, err := jwt.ParseEmailToken(token)
		if err != nil {
			code = e.ErrorAuthToken
		} else if time.Now().Unix() > claims.ExpiresAt {
			code = e.ErrorAuthCheckTokenTimeout
		} else {
			userId = claims.UserId
			email = claims.Email
			password = claims.Password
			operationType = claims.OperationType
		}
	}
	if code != e.SUCCESS {
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	//获取该用户的信息
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetUserById(userId)
	if err != nil {
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	if operationType == 1 {
		//绑定邮箱
		user.Email = email
	} else if operationType == 2 {
		//解绑邮箱
		user.Email = ""
	} else if operationType == 3 {
		//修改密码
		err = user.SetPassword(password)
		if err != nil {
			code = e.ERROR
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
			}
		}
	}
	err = userDao.UpdateUserById(userId, user)
	if err != nil {
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildUser(user),
	}
}
