// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package database

type Product struct {
	ID            int64
	Sku           string
	Name          string
	Description   string
	Price         float64
	Image         string
	StockQuantity int32
	TaxPercetange int32
}