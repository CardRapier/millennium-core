// internal/product/handler.go
package product

import (
	"database/sql"
	"fmt"
	database "millennium/db"
	"millennium/db/sqlc"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

func GetProductsHandler(c *fiber.Ctx) error {
	searchParam := c.Query("search")
	nullSearchParam := sql.NullString{String: searchParam, Valid: len(searchParam) > 0}
	fmt.Println("Search Parameter:", searchParam)

	var products []sqlc.Product
	var err error
	if nullSearchParam.Valid {
		products, err = database.Queries.GetFilteredProducts(c.Context(), nullSearchParam)
	} else {
		products, err = database.Queries.GetAllProducts(c.Context())
	}
	if err != nil {
		return err
	}
	fmt.Println(products)
	if products == nil {
		products = []sqlc.Product{}
	}
	return c.JSON(products)
}

func CreateProductsHandler(c *fiber.Ctx) error {
	var createProduct sqlc.AddProductParams
	if err := c.BodyParser(&createProduct); err != nil {
		fmt.Print(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": fmt.Sprintf("Invalid request format %d", err)})
	}
	id, err := database.Queries.AddProduct(c.Context(), createProduct)
	if err != nil {
		return c.JSON(fiber.Map{"message": "Error creating the product"});
	}
	return c.JSON(fiber.Map{"message": fmt.Sprintf("Product added successfully: %d", id)})
}

//TODO: Implement this with sqlc
func UpdateProductsHandler(c *fiber.Ctx) error {
	var updateProduct sqlc.UpdateProductParams
	if err := c.BodyParser(&updateProduct); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request format"})
	}
	err := database.Queries.UpdateProduct(c.Context(),updateProduct)
	if err != nil {
		return c.JSON(fiber.Map{"message": "Error creating the product"});
	}
	return c.JSON(fiber.Map{"message": "Product updated successfully"})
}
