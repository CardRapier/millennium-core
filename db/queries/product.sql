-- queries/product.sql

-- name: GetAllProducts :many
SELECT * FROM Product;

-- name: GetFilteredProducts :many
SELECT *
    FROM Product
    WHERE to_char(id, '9999999') LIKE '%' || $1 || '%'
        OR name ILIKE '%' || $1 || '%';
-- name: GetProductByID :one
SELECT * FROM Product WHERE id = $1;

-- name: GetProductsByID :many
SELECT * FROM Product WHERE id = ANY($1::int[]);

-- name: AddProduct :one
INSERT INTO Product (utc, name, price, iva, gross_margin) VALUES ($1, $2, $3, $4, $5) RETURNING id;

-- name: UpdateProduct :exec
UPDATE Product SET utc = $1, name = $2, price = $3, iva = $4, gross_margin = $5 WHERE id = $6 RETURNING *;

-- name: DeleteProduct :exec
DELETE FROM Product WHERE id = $1;