package repository

import (
	"context"

	"syaiful.com/simple-commerce/entity"
)

type UserRepository interface {
	Authentication(ctx context.Context, username string) (entity.User, error)
	Create(username string, password string, roles []string)
}
