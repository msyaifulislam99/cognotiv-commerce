package impl

import (
	"context"
	"errors"

	"gorm.io/gorm"
	"syaiful.com/simple-commerce/entity"
	"syaiful.com/simple-commerce/repository"
)

func NewOrderDetailRepositoryImpl(DB *gorm.DB) repository.OrderDetailRepository {
	return &orderDetailRepositoryImpl{DB: DB}
}

type orderDetailRepositoryImpl struct {
	*gorm.DB
}

func (orderDetailRepository *orderDetailRepositoryImpl) FindById(ctx context.Context, id string) (entity.OrderDetail, error) {
	var orderDetail entity.OrderDetail
	result := orderDetailRepository.DB.WithContext(ctx).Where("order_detail_id = ?", id).Preload("Product").First(&orderDetail)
	if result.RowsAffected == 0 {
		return entity.OrderDetail{}, errors.New("order Detail Not Found")
	}
	return orderDetail, nil
}
