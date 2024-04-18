package dao

import (
	"context"
	"gin-mall/repository/db/model"
	"gorm.io/gorm"
)

type FavoriteDao struct {
	*gorm.DB
}

func NewFavoriteDao(ctx context.Context) *FavoriteDao {
	return &FavoriteDao{NewDBClient(ctx)}
}

func NewFavoriteDaoByDB(db *gorm.DB) *FavoriteDao {
	return &FavoriteDao{db}
}

// FavoriteExistOrNot 判断是否存在
func (dao *FavoriteDao) FavoriteExistOrNot(pId, uId uint) (exist bool, err error) {
	var count int64

	err = dao.DB.Model(&model.Favorite{}).
		Where("product_id=? AND user_id=?", pId, uId).
		Count(&count).Error

	if count == 0 || err != nil {
		return false, err
	}
	return true, err

}

// CreateFavorite 创建收藏夹
func (dao *FavoriteDao) CreateFavorite(favorite *model.Favorite) (err error) {
	err = dao.DB.Create(&favorite).Error
	return
}

// ListFavoriteByUserId 通过 user_id 获取收藏夹列表
func (dao *FavoriteDao) ListFavoriteByUserId(uId uint, pageSize, pageNum int) (favorites []*model.Favorite, total int64, err error) {
	// 总数
	err = dao.DB.Model(&model.Favorite{}).Preload("User").
		Where("user_id=?", uId).Count(&total).Error
	if err != nil {
		return
	}
	// 分页
	err = dao.DB.Model(model.Favorite{}).Preload("User").Where("user_id=?", uId).
		Offset((pageNum - 1) * pageSize).
		Limit(pageSize).Find(&favorites).Error
	return
}

// DeleteFavoriteById 删除收藏夹
func (dao *FavoriteDao) DeleteFavoriteById(fId uint) error {
	return dao.DB.Where("id=?", fId).Delete(&model.Favorite{}).Error
}
