package service

import (
	"context"

	"syaiful.com/simple-commerce/entity"
	"syaiful.com/simple-commerce/model"
)

type UserService interface {
	Authentication(ctx context.Context, model model.UserModel) entity.User
}
