package controller

import (
	"github.com/gofiber/fiber/v2"
	"syaiful.com/simple-commerce/configuration"
	"syaiful.com/simple-commerce/middleware"
	"syaiful.com/simple-commerce/model"
	"syaiful.com/simple-commerce/service"
)

type OrderDetailController struct {
	service.OrderDetailService
	configuration.Config
}

func NewOrderDetailController(orderDetailService *service.OrderDetailService, config configuration.Config) *OrderDetailController {
	return &OrderDetailController{OrderDetailService: *orderDetailService, Config: config}
}

func (controller OrderDetailController) Route(app *fiber.App) {
	app.Get("/v1/api/order-detail/:id", middleware.AuthenticateJWT("ROLE_USER", controller.Config), controller.FindById)
}

func (controller OrderDetailController) FindById(c *fiber.Ctx) error {
	id := c.Params("id")

	result := controller.OrderDetailService.FindById(c.Context(), id)
	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    result,
	})
}
