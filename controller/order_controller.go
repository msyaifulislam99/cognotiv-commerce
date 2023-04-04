package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"syaiful.com/simple-commerce/configuration"
	"syaiful.com/simple-commerce/exception"
	"syaiful.com/simple-commerce/middleware"
	"syaiful.com/simple-commerce/model"
	"syaiful.com/simple-commerce/service"
)

type OrderController struct {
	service.OrderService
	configuration.Config
}

func NewOrderController(orderService *service.OrderService, config configuration.Config) *OrderController {
	return &OrderController{OrderService: *orderService, Config: config}
}

func (controller OrderController) Route(app *fiber.App) {
	app.Post("/v1/api/order", middleware.AuthenticateJWT("ROLE_USER", controller.Config), controller.Create)
	app.Delete("/v1/api/order/:id", middleware.AuthenticateJWT("ROLE_USER", controller.Config), controller.Delete)
	app.Get("/v1/api/order/:id", middleware.AuthenticateJWT("ROLE_USER", controller.Config), controller.FindById)
	app.Get("/v1/api/order", middleware.AuthenticateJWT("ROLE_USER", controller.Config), controller.FindAll)
}

func (controller OrderController) Create(c *fiber.Ctx) error {
	var request model.OrderCreateModel
	err := c.BodyParser(&request)
	exception.PanicLogging(err)

	idLocal := c.Locals("user_id")
	strId := fmt.Sprintf("%v", idLocal)

	response := controller.OrderService.Create(c.Context(), request, strId)
	return c.Status(fiber.StatusCreated).JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    response,
	})
}

// Delete func delete one exists order.
// @Description delete one exists order.
// @Summary delete one exists order
// @Tags Order
// @Accept json
// @Produce json
// @Param id path string true "Order Id"
// @Success 200 {object} model.GeneralResponse
// @Security JWT
// @Router /v1/api/order/{id} [delete]
func (controller OrderController) Delete(c *fiber.Ctx) error {
	id := c.Params("id")

	controller.OrderService.Delete(c.Context(), id)
	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
	})
}

// FindById func gets one exists order.
// @Description Get one exists order.
// @Summary get one exists order
// @Tags Order
// @Accept json
// @Produce json
// @Param id path string true "Order Id"
// @Success 200 {object} model.GeneralResponse
// @Security JWT
// @Router /v1/api/order/{id} [get]
func (controller OrderController) FindById(c *fiber.Ctx) error {
	id := c.Params("id")

	result := controller.OrderService.FindById(c.Context(), id)
	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    result,
	})
}

// FindAll func gets all exists order.
// @Description Get all exists order.
// @Summary get all exists order
// @Tags Order
// @Accept json
// @Produce json
// @Success 200 {object} model.GeneralResponse
// @Security JWT
// @Router /v1/api/order [get]
func (controller OrderController) FindAll(c *fiber.Ctx) error {
	idSession := c.Locals("user_id")
	strId := fmt.Sprintf("%v", idSession)

	result := controller.OrderService.FindAll(c.Context(), strId)
	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    result,
	})
}
