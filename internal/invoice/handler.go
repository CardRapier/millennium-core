package invoice

import (
	"database/sql"
	"fmt"
	"log"
	database "millennium/db"
	"millennium/db/sqlc"
	"time"

	"github.com/gofiber/fiber/v2"
)

type CreateSaleDTO struct {
	Products    map[int]Product `json:"products"`
	Date        int64 `json:"date"`
	// PaymentType int32     `json:"payment_type"`
	// UserID      int32     `json:"user_id"`
}

type Product struct {
	ID       int32 `json:"id"`
	Quantity int32 `json:"quantity"`
}

type Sales struct {
	ProductID int32 `json:"product_id"`
	InvoiceID int32 `json:"invoice_id"`
	Quantity  int32 `json:"quantity"`
	Price     int32 `json:"price"`
	Iva       int32 `json:"iva"`
	Subtotal  int32 `json:"subtotal"`
}

type Invoice struct {
	Date        time.Time `json:"date"`
	PaymentType int32     `json:"payment_type"`
	// UserID int32 `json:"user_id"`
}

func CreateSalesHandler(c *fiber.Ctx) error {
	var createSale CreateSaleDTO
	if err := c.BodyParser(&createSale); err != nil {
		fmt.Print(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": fmt.Sprintf("Invalid request format %d", err)})
	}

	var productIDs []int32
	productMap := make(map[int32]int32)
	for _, product := range createSale.Products {
		productMap[product.ID] = product.Quantity
		productIDs = append(productIDs, product.ID)
	}

	products, err := database.Queries.GetProductsByID(c.Context(), productIDs)
	if err != nil {
		fmt.Print(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Error getting the products"})
	}

	invoice := sqlc.AddInvoiceParams{
		Date:          createSale.Date,
		PaymentTypeID: 1,
		// UserID:
	}

	invoiceId, err := database.Queries.AddInvoice(c.Context(), invoice)
	if err != nil {
		fmt.Print(err)
		return c.JSON(fiber.Map{"message": "Error creating the invoice"})
	}

	var total sql.NullInt32
	total.Int32 = 0
	for _, product := range products {
		subtotal := product.Price * productMap[product.ID]
		sale := sqlc.AddSalesParams{
			ProductID: product.ID,
			InvoiceID: invoiceId,
			Quantity:  productMap[product.ID],
			Price:     product.Price,
			Iva:       product.Iva,
			Subtotal:  subtotal,
		}

		total.Int32 += subtotal
		total.Valid = true

		_, err := database.Queries.AddSales(c.Context(), sale)
		if err != nil {
			log.Fatal(err)
		}
	}

	updateInvoice := sqlc.UpdateInvoiceParams{
		ID: invoiceId,
		Total: total,
	}

	database.Queries.UpdateInvoice(c.Context(), updateInvoice)

	return c.JSON(fiber.Map{"message": fmt.Sprintf("Invoice added successfully: %d", invoiceId)})

}
