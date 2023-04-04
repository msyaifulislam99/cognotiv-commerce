package impl

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"syaiful.com/simple-commerce/entity"
	"syaiful.com/simple-commerce/exception"
	"syaiful.com/simple-commerce/repository"
)

func NewUserRepositoryImpl(DB *gorm.DB) repository.UserRepository {
	return &userRepositoryImpl{DB: DB}
}

type userRepositoryImpl struct {
	*gorm.DB
}

func (userRepository *userRepositoryImpl) Create(name string, password string, roles []string) {
	var userRoles []entity.UserRole
	id := uuid.New()
	for _, role := range roles {
		userRoles = append(userRoles, entity.UserRole{
			Role:   role,
			UserId: id,
		})
	}
	user := entity.User{
		Name:      name,
		Password:  password,
		IsActive:  true,
		UserRoles: userRoles,
	}
	err := userRepository.DB.Create(&user).Error
	exception.PanicLogging(err)
}

func (userRepository *userRepositoryImpl) DeleteAll() {
	err := userRepository.DB.Where("1=1").Delete(&entity.User{}).Error
	exception.PanicLogging(err)
}

func (userRepository *userRepositoryImpl) Authentication(ctx context.Context, email string) (entity.User, error) {
	var userResult entity.User
	result := userRepository.DB.WithContext(ctx).
		Joins(`inner join "user_role" on "user_role".user_id = "user".id`).
		Preload("UserRoles").
		Where(`"user".email = ? and "user".is_active = ?`, email, true).
		Find(&userResult)
	if result.RowsAffected == 0 {
		return entity.User{}, errors.New("user not found")
	}
	return userResult, nil
}
