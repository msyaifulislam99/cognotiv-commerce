package impl

import (
	"context"

	"github.com/google/uuid"
	"syaiful.com/simple-commerce/common"
	"syaiful.com/simple-commerce/entity"
	"syaiful.com/simple-commerce/exception"
	"syaiful.com/simple-commerce/model"
	"syaiful.com/simple-commerce/repository"
	"syaiful.com/simple-commerce/service"
)

func NewOrderServiceImpl(orderRepository *repository.OrderRepository) service.OrderService {
	return &orderServiceImpl{OrderRepository: *orderRepository}
}

type orderServiceImpl struct {
	repository.OrderRepository
}

func (orderService *orderServiceImpl) Create(ctx context.Context, orderModel model.OrderCreateUpdateModel) model.OrderCreateUpdateModel {
	common.Validate(orderModel)
	uuidGenerate := uuid.New()
	var orderDetails []entity.OrderDetail
	var totalPrice int64 = 0

	for _, detail := range orderModel.OrderDetails {
		totalPrice = totalPrice + detail.SubTotalPrice
		orderDetails = append(orderDetails, entity.OrderDetail{
			OrderId:       uuidGenerate,
			ProductId:     detail.ProductId,
			Id:            uuid.New(),
			SubTotalPrice: detail.SubTotalPrice,
			Price:         detail.Price,
			Quantity:      detail.Quantity,
		})
	}

	order := entity.Order{
		Id:           uuid.New(),
		TotalPrice:   totalPrice,
		OrderDetails: orderDetails,
	}

	orderService.OrderRepository.Insert(ctx, order)
	return orderModel
}

func (orderService *orderServiceImpl) Delete(ctx context.Context, id string) {
	order, err := orderService.OrderRepository.FindById(ctx, id)
	if err != nil {
		panic(exception.NotFoundError{
			Message: err.Error(),
		})
	}
	orderService.OrderRepository.Delete(ctx, order)
}

func (orderService *orderServiceImpl) FindById(ctx context.Context, id string) model.OrderModel {
	order, err := orderService.OrderRepository.FindById(ctx, id)
	if err != nil {
		panic(exception.NotFoundError{
			Message: err.Error(),
		})
	}
	var orderDetails []model.OrderDetailModel
	for _, detail := range order.OrderDetails {
		orderDetails = append(orderDetails, model.OrderDetailModel{
			Id:            detail.Id.String(),
			SubTotalPrice: detail.SubTotalPrice,
			Price:         detail.Price,
			Quantity:      detail.Quantity,
			Product: model.ProductModel{
				Id:          detail.Product.Id.String(),
				Name:        detail.Product.Name,
				Price:       detail.Product.Price,
				Description: detail.Product.Description,
				Image:       detail.Product.Image,
			},
		})
	}

	return model.OrderModel{
		Id:           order.Id.String(),
		TotalPrice:   order.TotalPrice,
		OrderDetails: orderDetails,
	}
}

func (orderService *orderServiceImpl) FindAll(ctx context.Context) (responses []model.OrderModel) {
	orders := orderService.OrderRepository.FindAll(ctx)
	for _, order := range orders {
		var orderDetails []model.OrderDetailModel
		for _, detail := range order.OrderDetails {
			orderDetails = append(orderDetails, model.OrderDetailModel{
				Id:            detail.Id.String(),
				SubTotalPrice: detail.SubTotalPrice,
				Price:         detail.Price,
				Quantity:      detail.Quantity,
				Product: model.ProductModel{
					Id:          detail.Product.Id.String(),
					Name:        detail.Product.Name,
					Price:       detail.Product.Price,
					Description: detail.Product.Description,
					Image:       detail.Product.Image,
				},
			})
		}

		responses = append(responses, model.OrderModel{
			Id:           order.Id.String(),
			TotalPrice:   order.TotalPrice,
			OrderDetails: orderDetails,
		})
	}

	return responses
}
