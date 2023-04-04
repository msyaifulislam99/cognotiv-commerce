package main

import (
	"github.com/gofiber/fiber/v2"
	"syaiful.com/simple-commerce/configuration"
	"syaiful.com/simple-commerce/controller"
	"syaiful.com/simple-commerce/exception"
	repository "syaiful.com/simple-commerce/repository/impl"
	service "syaiful.com/simple-commerce/service/impl"
)

func main() {
	config := configuration.New()
	database := configuration.NewDatabase(config)

	// repo
	productRepository := repository.NewProductRepositoryImpl(database)
	userRepository := repository.NewUserRepositoryImpl(database)
	orderRepository := repository.NewOrderRepositoryImpl(database)
	orderDetailRepository := repository.NewOrderDetailRepositoryImpl(database)

	productService := service.NewProductServiceImpl(&productRepository)
	userService := service.NewUserServiceImpl(&userRepository)
	orderService := service.NewOrderServiceImpl(&orderRepository)
	orderDetailService := service.NewOrderDetailServiceImpl(&orderDetailRepository)

	productController := controller.NewProductController(&productService, config)
	userController := controller.NewUserController(&userService, config)
	orderController := controller.NewOrderController(&orderService, config)
	orderDetailController := controller.NewOrderDetailController(&orderDetailService, config)

	app := fiber.New(fiber.Config{
		ErrorHandler: exception.ErrorHandler,
	})

	productController.Route(app)
	userController.Route(app)
	orderController.Route(app)
	orderDetailController.Route(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 2!")
	})

	app.Listen(":8000")
}
