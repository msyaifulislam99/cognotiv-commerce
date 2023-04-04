package repository

import (
	"context"

	"syaiful.com/simple-commerce/entity"
)

type OrderRepository interface {
	Insert(ctx context.Context, transaction entity.Order) entity.Order
	Delete(ctx context.Context, transaction entity.Order)
	FindById(ctx context.Context, id string) (entity.Order, error)
	FindAll(ctx context.Context) []entity.Order
	FindAllPending(ctx context.Context) []entity.Order
	FindMyOrders(ctx context.Context, userId string) []entity.Order
}
