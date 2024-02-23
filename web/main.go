// web/main.go
package web

import (
	"fmt"
	"millennium/internal/invoice"
	"millennium/internal/product"

	"github.com/gofiber/fiber/v2"
)

func InitializeServer() {
    app := fiber.New()

    // Define API endpoints and handlers
    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Welcome to the POS system!")
    })

    app.Get("/products", product.GetProductsHandler)
    app.Post("/products", product.CreateProductsHandler)
    app.Put("/products", product.UpdateProductsHandler)
    app.Post("/invoice", invoice.CreateSalesHandler)

    port := 8080
    fmt.Printf("Web server listening on port %d...\n", port)
    app.Listen(fmt.Sprintf(":%d", port))
}
