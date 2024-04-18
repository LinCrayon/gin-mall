package dao

import (
	"context"
	"gin-mall/repository/db/model"
	"gorm.io/gorm"
)

type CarouselDao struct {
	*gorm.DB
}

// NewCarouselDao /* 在内部负责创建一个新的数据库连接对象，并使用传入的上下文信息进行初始化。*/
func NewCarouselDao(ctx context.Context) *CarouselDao {
	return &CarouselDao{NewDBClient(ctx)}
}

// NewCarouselDaoByDB /*不负责创建数据库连接对象，而是依赖外部传入的已存在的数据库连接对象。*/
func NewCarouselDaoByDB(db *gorm.DB) *CarouselDao {
	return &CarouselDao{db}
}

// ListCarousel GetCarouselById 根据id获取 Carousel
func (dao *CarouselDao) ListCarousel() (r []model.Carousel, err error) {
	err = dao.DB.Model(&model.Carousel{}).Find(&r).Error
	return
}
