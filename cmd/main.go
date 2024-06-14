package main

import (
	"millennium/internal/services/products"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	productsHttpHandler := products.NewProductsHttpHandler()

	api := app.Group("/api")
	v1 := api.Group("/v1")

	productsHttpHandler.RegisterRoutes(v1)
}