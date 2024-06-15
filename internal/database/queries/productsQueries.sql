-- name: GetProducts :many
SELECT * FROM products;

-- name: CreateProduct :exec
INSERT INTO products (
    sku,
    name,
    description,
    price,
    image,
    stock_quantity,
    tax_percetange
) VALUES ($1, $2, $3, $4, $5, $6, $7);