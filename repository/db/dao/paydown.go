package dao

import (
	"context"
	"gorm.io/gorm"
)

type PayDownDao struct {
	*gorm.DB
}

func NewPayDownDao(ctx context.Context) *PayDownDao {
	return &PayDownDao{NewDBClient(ctx)}
}

func NewPayDownDaoByDB(db *gorm.DB) *PayDownDao {
	return &PayDownDao{db}
}
