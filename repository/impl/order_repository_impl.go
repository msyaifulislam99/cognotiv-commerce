package impl

import (
	"context"
	"errors"

	"gorm.io/gorm"
	"syaiful.com/simple-commerce/entity"
	"syaiful.com/simple-commerce/exception"
	"syaiful.com/simple-commerce/repository"
)

func NewOrderRepositoryImpl(DB *gorm.DB) repository.OrderRepository {
	return &orderRepositoryImpl{DB: DB}
}

type orderRepositoryImpl struct {
	*gorm.DB
}

func (orderRepository *orderRepositoryImpl) Insert(ctx context.Context, order entity.Order) entity.Order {
	err := orderRepository.DB.WithContext(ctx).Create(&order).Error
	exception.PanicLogging(err)
	return order
}

func (orderRepository *orderRepositoryImpl) Delete(ctx context.Context, order entity.Order) {
	orderRepository.DB.WithContext(ctx).Delete(&order)
}

func (orderRepository *orderRepositoryImpl) FindById(ctx context.Context, id string) (entity.Order, error) {
	var order entity.Order
	result := orderRepository.DB.WithContext(ctx).
		Table("order").
		Select("order.order_id, order.total_price, order_detail.order_detail_id, order_detail.sub_total_price, order_detail.price, order_detail.quantity, tb_product.product_id, tb_product.name, tb_product.price, tb_product.quantity").
		Joins("join order_detail on order_detail.order_id = order.order_id").
		Joins("join tb_product on tb_product.product_id = order_detail.product_id").
		Preload("OrderDetails").
		Preload("OrderDetails.Product").
		Where("order.order_id = ?", id).
		First(&order)
	if result.RowsAffected == 0 {
		return entity.Order{}, errors.New("order Not Found")
	}
	return order, nil
}

func (orderRepository *orderRepositoryImpl) FindAll(ctx context.Context) []entity.Order {
	var orders []entity.Order
	orderRepository.DB.WithContext(ctx).
		Table("order").
		Select(`"order".*`).
		Preload("OrderDetails").
		Preload("OrderDetails.Product").
		Preload("User").
		Find(&orders)
	return orders
}

func (orderRepository *orderRepositoryImpl) FindMyOrders(ctx context.Context, userId string) []entity.Order {
	var orders []entity.Order
	orderRepository.DB.WithContext(ctx).
		Table("order").
		Select(`"order".*`).
		Preload("OrderDetails").
		Preload("OrderDetails.Product").
		Where(`"order".user_id = ?`, userId).
		Find(&orders)
	return orders
}
