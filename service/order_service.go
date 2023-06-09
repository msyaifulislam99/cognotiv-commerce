package service

import (
	"context"

	"syaiful.com/simple-commerce/model"
)

type OrderService interface {
	Create(ctx context.Context, model model.OrderCreateModel, userId string) model.OrderCreateUpdateModel
	Delete(ctx context.Context, id string)
	FindById(ctx context.Context, id string) model.OrderModel
	FindAll(ctx context.Context) []model.OrderModel
	NotifyUser(ctx context.Context) error
	FindAllWithUser(ctx context.Context) []model.OrderModelWithUser
	FindMyOrder(ctx context.Context, userId string) []model.OrderModel
}
