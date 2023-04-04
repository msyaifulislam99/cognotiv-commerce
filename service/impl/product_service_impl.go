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

func NewProductServiceImpl(productRepository *repository.ProductRepository) service.ProductService {
	return &productServiceImpl{repo: *productRepository}
}

type productServiceImpl struct {
	repo repository.ProductRepository
}

func (service *productServiceImpl) Create(ctx context.Context, productModel model.ProductCreateOrUpdateModel) model.ProductCreateOrUpdateModel {
	common.Validate(productModel)
	product := entity.Product{
		Name:        productModel.Name,
		Price:       productModel.Price,
		Description: productModel.Description,
		Image:       productModel.Image,
	}
	service.repo.Insert(ctx, product)
	return productModel
}

func (service *productServiceImpl) Update(ctx context.Context, productModel model.ProductCreateOrUpdateModel, id string) model.ProductCreateOrUpdateModel {
	common.Validate(productModel)
	product := entity.Product{
		Id:          uuid.MustParse(id),
		Name:        productModel.Name,
		Price:       productModel.Price,
		Description: productModel.Description,
		Image:       productModel.Image,
	}
	service.repo.Update(ctx, product)
	return productModel
}

func (service *productServiceImpl) Delete(ctx context.Context, id string) {
	product, err := service.repo.FindById(ctx, id)
	if err != nil {
		panic(exception.NotFoundError{
			Message: err.Error(),
		})
	}
	service.repo.Delete(ctx, product)
}

func (service *productServiceImpl) FindById(ctx context.Context, id string) model.ProductModel {
	product, _ := service.repo.FindById(ctx, id)
	return model.ProductModel{
		Id:          product.Id.String(),
		Name:        product.Name,
		Price:       product.Price,
		Description: product.Description,
		Image:       product.Image,
	}
}

func (service *productServiceImpl) FindAll(ctx context.Context) (responses []model.ProductModel) {
	products := service.repo.FindAl(ctx)
	for _, product := range products {
		responses = append(responses, model.ProductModel{
			Id:          product.Id.String(),
			Name:        product.Name,
			Price:       product.Price,
			Description: product.Description,
			Image:       product.Image,
		})
	}
	if len(products) == 0 {
		return []model.ProductModel{}
	}
	return responses
}
