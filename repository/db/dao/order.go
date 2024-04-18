package dao

import (
	"context"
	"gin-mall/repository/db/model"
	"gorm.io/gorm"
)

type OrderDao struct {
	*gorm.DB
}

// NewOrderDao /* 在内部负责创建一个新的数据库连接对象，并使用传入的上下文信息进行初始化。*/
func NewOrderDao(ctx context.Context) *OrderDao {
	return &OrderDao{NewDBClient(ctx)}
}

// NewOrderDaoByDB /*不负责创建数据库连接对象，而是依赖外部传入的已存在的数据库连接对象。*/
func NewOrderDaoByDB(db *gorm.DB) *OrderDao {
	return &OrderDao{db}
}

// CreateOrder 创建订单
func (dao *OrderDao) CreateOrder(order *model.Order) error {
	return dao.DB.Create(&order).Error
}

// ListOrderByCondition 获取订单List
func (dao *OrderDao) ListOrderByCondition(condition map[string]interface{}, page model.BasePage) (orders []*model.Order, total int64, err error) {
	err = dao.DB.Model(&model.Order{}).Where(condition).
		Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = dao.DB.Model(&model.Order{}).Where(condition).
		Offset((page.PageNum - 1) * page.PageSize).
		Limit(page.PageSize).Order("created_at desc").Find(&orders).Error
	return
}

// GetOrderById 获取订单详情
func (dao *OrderDao) GetOrderById(id uint) (order *model.Order, err error) {
	err = dao.DB.Model(&model.Order{}).Where("id=?", id).
		First(&order).Error
	return
}

// DeleteOrderById 获取订单详情
/*func (dao *OrderDao) DeleteOrderById(id uint) error {
	return dao.DB.Where("id=?", id).Delete(&model.Order{}).Error
}*/
func (dao *OrderDao) DeleteOrderById(id uint) error {
	return dao.DB.Model(&model.Order{}).
		Where("id=?", id).
		Delete(&model.Order{}).Error
}

// UpdateOrderById 更新订单详情
func (dao *OrderDao) UpdateOrderById(id uint, order *model.Order) error {
	return dao.DB.Where("id=?", id).Updates(order).Error
}
