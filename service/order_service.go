package service

import (
	"context"

	"syaiful.com/simple-commerce/model"
)

type OrderService interface {
	Create(ctx context.Context, model model.OrderCreateModel) model.OrderCreateUpdateModel
	Delete(ctx context.Context, id string)
	FindById(ctx context.Context, id string) model.OrderModel
	FindAll(ctx context.Context) []model.OrderModel
}
