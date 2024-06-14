package products

import "github.com/gofiber/fiber/v2"

type ProductsHttpHandler struct {}

func NewProductsHttpHandler() *ProductsHttpHandler {
	return &ProductsHttpHandler{}
}

func (p *ProductsHttpHandler) RegisterRoutes(router fiber.Router) {
	productsRoute := router.Group("/products")
	productsRoute.Get("/", p.getProducts)
}

func (p *ProductsHttpHandler) getProducts(c *fiber.Ctx) error {
	return nil
}