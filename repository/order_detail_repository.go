package repository

import (
	"context"

	"syaiful.com/simple-commerce/entity"
)

type OrderDetailRepository interface {
	FindById(ctx context.Context, id string) (entity.OrderDetail, error)
}
