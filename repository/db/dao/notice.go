package dao

import (
	"context"
	"gin-mall/repository/db/model"
	"gorm.io/gorm"
)

type NoticeDao struct {
	*gorm.DB
}

// NewNoticeDao /* 在内部负责创建一个新的数据库连接对象，并使用传入的上下文信息进行初始化。*/
func NewNoticeDao(ctx context.Context) *NoticeDao {
	return &NoticeDao{NewDBClient(ctx)}
}

// NewNoticeDaoByDB /*不负责创建数据库连接对象，而是依赖外部传入的已存在的数据库连接对象。*/
func NewNoticeDaoByDB(db *gorm.DB) *NoticeDao {
	return &NoticeDao{db}
}

// GetNoticeById 根据id获取 notice
func (dao *NoticeDao) GetNoticeById(id uint) (notice *model.Notice, err error) {
	err = dao.DB.Model(&model.Notice{}).Where("id=?", id).First(&notice).Error
	return
}
