package controller

import (
	"github.com/gofiber/fiber/v2"
	"syaiful.com/simple-commerce/model"
	"syaiful.com/simple-commerce/service"
)

type HttpBinController struct {
	service.HttpBinService
}

func NewHttpBinController(httpBinService *service.HttpBinService) *HttpBinController {
	return &HttpBinController{HttpBinService: *httpBinService}
}

func (controller HttpBinController) Route(app *fiber.App) {
	app.Get("/v1/api/httpbin", controller.PostHttpBin)
}

func (controller HttpBinController) PostHttpBin(c *fiber.Ctx) error {

	controller.HttpBinService.PostMethod(c.Context())
	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    nil,
	})
}
