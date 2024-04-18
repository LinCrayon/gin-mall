package dao

import (
	"context"
	"gin-mall/repository/db/model"
	"gorm.io/gorm"
)

type UserDao struct {
	*gorm.DB
}

// NewUserDao /* 在内部负责创建一个新的数据库连接对象，并使用传入的上下文信息进行初始化。*/
func NewUserDao(ctx context.Context) *UserDao {
	return &UserDao{NewDBClient(ctx)}
}

// NewUserDaoByDB /*不负责创建数据库连接对象，而是依赖外部传入的已存在的数据库连接对象。*/
func NewUserDaoByDB(db *gorm.DB) *UserDao {
	return &UserDao{db}
}

// ExistOrNotByUserName 根据username判断是否存在该名字
func (dao *UserDao) ExistOrNotByUserName(userName string) (user *model.User, exist bool, err error) {
	var count int64
	//统计符合条件的记录数量
	err = dao.DB.Model(&model.User{}).Where("user_name = ?", userName).Count(&count).Error
	if count == 0 {
		return user, false, err
	}
	// 记录数不为零，则查询第一条符合条件的记录
	err = dao.DB.Model(&model.User{}).Where("user_name = ?", userName).First(&user).Error
	if err != nil {
		return user, false, err
	}
	return user, true, nil
}

// CreateUser 创建用户
func (dao *UserDao) CreateUser(user *model.User) (err error) {
	return dao.DB.Model(&model.User{}).Create(&user).Error
}

// GetUserById 根据id获取user
func (dao *UserDao) GetUserById(id uint) (user *model.User, err error) {
	err = dao.DB.Model(&model.User{}).Where("id=?", id).First(&user).Error
	return
}

// UpdateUserById 根据uid更新user信息
func (dao *UserDao) UpdateUserById(uid uint, user *model.User) error {
	//Update()需要手动指定更新的字段和值 而 Updates()会自动更新
	return dao.DB.Model(&model.User{}).Where("id=?", uid).Updates(&user).Error
	//return dao.DB.Model(&model.User{}).Where("id=?", uid).Update("字段名", "更新值").Error
}
