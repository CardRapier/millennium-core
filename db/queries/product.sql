-- queries/product.sql

-- name: GetAllProducts :many
SELECT * FROM Product;

-- name: GetProductByID :one
SELECT * FROM Product WHERE id = $1;

-- name: GetProductsByID :many
SELECT * FROM Product WHERE id = ANY($1::int[]);

-- name: AddProduct :one
INSERT INTO Product (name, price, iva, gross_margin) VALUES ($1, $2, $3, $4) RETURNING id;

-- name: UpdateProduct :exec
UPDATE Product SET name = $1, price = $2, iva = $3, gross_margin = $4 WHERE id = $5 RETURNING *;

-- name: DeleteProduct :exec
DELETE FROM Product WHERE id = $1;