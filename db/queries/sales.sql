-- queries/sales.sql

-- name: AddSales :one
INSERT INTO Sales (product_id, invoice_id, quantity, price, iva, subtotal) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id;