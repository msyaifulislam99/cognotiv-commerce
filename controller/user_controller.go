package controller

import (
	"github.com/gofiber/fiber/v2"
	"syaiful.com/simple-commerce/common"
	"syaiful.com/simple-commerce/configuration"
	"syaiful.com/simple-commerce/exception"
	"syaiful.com/simple-commerce/model"
	"syaiful.com/simple-commerce/service"
)

func NewUserController(userService *service.UserService, config configuration.Config) *UserController {
	return &UserController{UserService: *userService, Config: config}
}

type UserController struct {
	service.UserService
	configuration.Config
}

func (controller UserController) Route(app *fiber.App) {
	app.Post("/v1/api/authentication", controller.Authentication)
}

func (controller UserController) Authentication(c *fiber.Ctx) error {
	var request model.UserModel
	err := c.BodyParser(&request)
	exception.PanicLogging(err)

	result := controller.UserService.Authentication(c.Context(), request)
	var userRoles []map[string]interface{}
	for _, userRole := range result.UserRoles {
		userRoles = append(userRoles, map[string]interface{}{
			"role": userRole.Role,
		})
	}
	tokenJwtResult := common.GenerateToken(result.Email, userRoles, controller.Config)
	resultWithToken := map[string]interface{}{
		"token": tokenJwtResult,
		"email": result.Email,
		"role":  userRoles,
	}
	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    resultWithToken,
	})
}
