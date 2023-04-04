package service

import (
	"context"

	"syaiful.com/simple-commerce/model"
)

type OrderDetailService interface {
	FindById(ctx context.Context, id string) model.OrderDetailModel
}
