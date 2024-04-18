package serializer

import (
	"gin-mall/config"
	"gin-mall/repository/db/model"
)

type User struct {
	Id       uint   `json:"id"`
	UserName string `json:"user_name"`
	NickName string `json:"nickname"`
	Type     int    `json:"type"`
	Email    string `json:"email"`
	Status   string `json:"status"`
	Avatar   string `json:"avatar"`
	CreateAt int64  `json:"create_at"`
}

// BuildUser 序列化用户
func BuildUser(user *model.User) *User {
	return &User{
		Id:       user.ID,
		UserName: user.UserName,
		NickName: user.NickName,
		Email:    user.Email,
		Status:   user.Status,
		Avatar:   config.Host + config.HttpPort + config.AvatarPath + user.Avatar,
		CreateAt: user.CreatedAt.Unix(),
	}
}
