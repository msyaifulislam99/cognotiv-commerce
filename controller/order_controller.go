package controller

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

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
	app.Get("/v1/api/order", middleware.AuthenticateJWT("ROLE_USER", controller.Config), controller.FindMyOrderAll)
	app.Get("/v1/api/all-order", middleware.AuthenticateJWT("ROLE_ADMIN", controller.Config), controller.FindAll)
	app.Get("/v1/api/get-report", middleware.AuthenticateJWT("ROLE_ADMIN", controller.Config), controller.GenerateReport)
}

func (controller OrderController) GenerateReport(c *fiber.Ctx) error {
	file, err := os.Create("report.csv")
	if err != nil {
		exception.PanicLogging(err)
	}
	defer file.Close()

	// Write the CSV data
	writer := csv.NewWriter(file)

	header := []string{"OrderID", "CustomerName", "OrderDate", "TotalPrice", "Status"}
	writer.Write(header)

	data := controller.OrderService.FindAllWithUser(c.Context())

	for _, row := range data {
		writer.Write([]string{row.Id, *row.User.Name, row.TransactionDate.String(), strconv.FormatInt(row.TotalPrice, 10), row.Status})
	}
	defer writer.Flush()

	// Set the response headers
	c.Set("Content-Type", "text/csv")
	c.Set("Content-Disposition", "attachment; filename=report.csv")

	// Send the file for download
	return c.SendFile("report.csv")
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

func (controller OrderController) Delete(c *fiber.Ctx) error {
	id := c.Params("id")

	controller.OrderService.Delete(c.Context(), id)
	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
	})
}
func (controller OrderController) FindById(c *fiber.Ctx) error {
	id := c.Params("id")

	result := controller.OrderService.FindById(c.Context(), id)
	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    result,
	})
}
func (controller OrderController) FindAll(c *fiber.Ctx) error {
	result := controller.OrderService.FindAll(c.Context())
	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    result,
	})
}
func (controller OrderController) FindMyOrderAll(c *fiber.Ctx) error {
	idSession := c.Locals("user_id")
	strId := fmt.Sprintf("%v", idSession)

	result := controller.OrderService.FindMyOrder(c.Context(), strId)
	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    result,
	})
}
