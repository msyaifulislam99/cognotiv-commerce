package impl

import (
	"context"
	"time"

	"github.com/google/uuid"
	"syaiful.com/simple-commerce/common"
	"syaiful.com/simple-commerce/entity"
	"syaiful.com/simple-commerce/exception"
	"syaiful.com/simple-commerce/model"
	"syaiful.com/simple-commerce/repository"
	"syaiful.com/simple-commerce/service"
)

func NewOrderServiceImpl(orderRepository *repository.OrderRepository, _productRepo *repository.ProductRepository) service.OrderService {
	return &orderServiceImpl{repo: *orderRepository, productRepo: *_productRepo}
}

type orderServiceImpl struct {
	repo        repository.OrderRepository
	productRepo repository.ProductRepository
}

func (orderService *orderServiceImpl) Create(ctx context.Context, orderModel model.OrderCreateModel, userId string) model.OrderCreateUpdateModel {
	common.Validate(orderModel)
	uuidGenerate := uuid.New()
	var orderDetails []entity.OrderDetail
	var totalPrice int64 = 0

	for _, detail := range orderModel.Product {
		product, _ := orderService.productRepo.FindById(ctx, detail.Id)
		totalPrice = totalPrice + (product.Price * detail.Qty)
		orderDetails = append(orderDetails, entity.OrderDetail{
			OrderId:       uuidGenerate,
			ProductId:     product.Id,
			Id:            uuid.New(),
			SubTotalPrice: product.Price * detail.Qty,
			Price:         product.Price,
			Quantity:      int32(detail.Qty),
		})
	}

	orderId := uuid.New()
	parseUserId, err := uuid.Parse(userId)
	if err != nil {
		panic(exception.NotFoundError{
			Message: err.Error(),
		})
	}

	order := entity.Order{
		Id:              orderId,
		TotalPrice:      totalPrice,
		OrderDetails:    orderDetails,
		TransactionDate: time.Now(),
		UserId:          parseUserId,
	}

	orderService.repo.Insert(ctx, order)
	return model.OrderCreateUpdateModel{
		Id:         orderId.String(),
		TotalPrice: &order.TotalPrice,
	}
}

func (orderService *orderServiceImpl) Delete(ctx context.Context, id string) {
	order, err := orderService.repo.FindById(ctx, id)
	if err != nil {
		panic(exception.NotFoundError{
			Message: err.Error(),
		})
	}
	orderService.repo.Delete(ctx, order)
}

func (orderService *orderServiceImpl) FindById(ctx context.Context, id string) model.OrderModel {
	order, err := orderService.repo.FindById(ctx, id)
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

func (orderService *orderServiceImpl) FindAll(ctx context.Context, userId string) (responses []model.OrderModel) {
	orders := orderService.repo.FindMyOrders(ctx, userId)
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
