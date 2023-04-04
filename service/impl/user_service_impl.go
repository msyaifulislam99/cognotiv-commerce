package impl

import (
	"context"

	// "golang.org/x/crypto/bcrypt"
	"syaiful.com/simple-commerce/entity"
	"syaiful.com/simple-commerce/exception"
	"syaiful.com/simple-commerce/model"
	"syaiful.com/simple-commerce/repository"
	"syaiful.com/simple-commerce/service"
)

func NewUserServiceImpl(userRepository *repository.UserRepository) service.UserService {
	return &userServiceImpl{UserRepository: *userRepository}
}

type userServiceImpl struct {
	repository.UserRepository
}

func (userService *userServiceImpl) Authentication(ctx context.Context, model model.UserModel) entity.User {
	userResult, err := userService.UserRepository.Authentication(ctx, model.Email)
	if err != nil {
		panic(exception.UnauthorizedError{
			Message: err.Error(),
		})
	}
	// err = bcrypt.CompareHashAndPassword([]byte(userResult.Password), []byte(model.Password))
	// if err != nil {
	// 	panic(exception.UnauthorizedError{
	// 		Message: "incorrect email and password",
	// 	})
	// }
	return userResult
}
