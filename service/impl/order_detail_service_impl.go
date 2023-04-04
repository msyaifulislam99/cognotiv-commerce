package impl

import (
	"context"

	"syaiful.com/simple-commerce/exception"
	"syaiful.com/simple-commerce/model"
	"syaiful.com/simple-commerce/repository"
	"syaiful.com/simple-commerce/service"
)

func NewOrderDetailServiceImpl(orderDetailRepository *repository.OrderDetailRepository) service.OrderDetailService {
	return &orderDetailServiceImpl{OrderDetailRepository: *orderDetailRepository}
}

type orderDetailServiceImpl struct {
	repository.OrderDetailRepository
}

func (orderDetailService *orderDetailServiceImpl) FindById(ctx context.Context, id string) model.OrderDetailModel {
	orderDetail, err := orderDetailService.OrderDetailRepository.FindById(ctx, id)
	if err != nil {
		panic(exception.NotFoundError{
			Message: err.Error(),
		})
	}
	return model.OrderDetailModel{
		Id:            orderDetail.Id.String(),
		SubTotalPrice: orderDetail.SubTotalPrice,
		Price:         orderDetail.Price,
		Quantity:      orderDetail.Quantity,
		Product: model.ProductModel{
			Id:          orderDetail.Product.Id.String(),
			Name:        orderDetail.Product.Name,
			Price:       orderDetail.Product.Price,
			Description: orderDetail.Product.Description,
			Image:       orderDetail.Product.Image,
		},
	}
}
