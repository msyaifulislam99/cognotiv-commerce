package controller

import (
	"github.com/gofiber/fiber/v2"
	"syaiful.com/simple-commerce/configuration"
	"syaiful.com/simple-commerce/exception"
	"syaiful.com/simple-commerce/middleware"
	"syaiful.com/simple-commerce/model"
	"syaiful.com/simple-commerce/service"
)

type ProductController struct {
	service.ProductService
	configuration.Config
}

func NewProductController(productService *service.ProductService, config configuration.Config) *ProductController {
	return &ProductController{ProductService: *productService, Config: config}
}

func (controller ProductController) Route(app *fiber.App) {
	app.Post("/v1/api/product", middleware.AuthenticateJWT("ROLE_ADMIN", controller.Config), controller.Create)
	app.Put("/v1/api/product/:id", middleware.AuthenticateJWT("ROLE_ADMIN", controller.Config), controller.Update)
	app.Delete("/v1/api/product/:id", middleware.AuthenticateJWT("ROLE_ADMIN", controller.Config), controller.Delete)
	app.Get("/v1/api/product/:id", middleware.AuthenticateJWT("ROLE_ADMIN", controller.Config), controller.FindById)
	app.Get("/v1/api/product", controller.FindAll)
}

func (controller ProductController) Create(c *fiber.Ctx) error {
	var request model.ProductCreateOrUpdateModel
	err := c.BodyParser(&request)
	exception.PanicLogging(err)

	response := controller.ProductService.Create(c.Context(), request)
	return c.Status(fiber.StatusCreated).JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    response,
	})
}

func (controller ProductController) Update(c *fiber.Ctx) error {
	var request model.ProductCreateOrUpdateModel
	id := c.Params("id")
	err := c.BodyParser(&request)
	exception.PanicLogging(err)

	response := controller.ProductService.Update(c.Context(), request, id)
	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    response,
	})
}

func (controller ProductController) Delete(c *fiber.Ctx) error {
	id := c.Params("id")

	controller.ProductService.Delete(c.Context(), id)
	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
	})
}

func (controller ProductController) FindById(c *fiber.Ctx) error {
	id := c.Params("id")

	result := controller.ProductService.FindById(c.Context(), id)
	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    result,
	})
}

func (controller ProductController) FindAll(c *fiber.Ctx) error {
	result := controller.ProductService.FindAll(c.Context())
	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    result,
	})
}
