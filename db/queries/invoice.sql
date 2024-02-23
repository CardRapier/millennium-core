-- queries/invoice.sql

-- name: AddInvoice :one
INSERT INTO Invoice (date, payment_type_id) VALUES ($1, $2) RETURNING id;

-- name: UpdateInvoice :one
UPDATE Invoice SET total = $1 WHERE id = $2 RETURNING id;