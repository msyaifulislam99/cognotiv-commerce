package main

import (
	"context"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
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
	orderService := service.NewOrderServiceImpl(&orderRepository, &productRepository, config)
	orderDetailService := service.NewOrderDetailServiceImpl(&orderDetailRepository)

	productController := controller.NewProductController(&productService, config)
	userController := controller.NewUserController(&userService, config)
	orderController := controller.NewOrderController(&orderService, config)
	orderDetailController := controller.NewOrderDetailController(&orderDetailService, config)

	app := fiber.New(fiber.Config{
		ErrorHandler: exception.ErrorHandler,
	})

	app.Use(limiter.New(limiter.Config{
		Max:        100,
		Expiration: 1 * time.Minute,
	}))

	productController.Route(app)
	userController.Route(app)
	orderController.Route(app)
	orderDetailController.Route(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 2!")
	})

	go func() {
		for {
			// now := time.Now()

			// // Calculate duration until midnight
			// midnight := time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, now.Location())
			// duration := midnight.Sub(now)

			// // Sleep until midnight
			// time.Sleep(duration)
			time.Sleep(5 * time.Second)

			fmt.Println("background task running...")
			orderService.NotifyUser(context.Background())
			fmt.Println("background task done...")
		}
	}()

	app.Listen(":8000")
}
