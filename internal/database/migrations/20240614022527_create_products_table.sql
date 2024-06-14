-- +goose Up
CREATE TABLE products (
    id BIGSERIAL PRIMARY KEY,
    sku TEXT NOT NULL,
    name VARCHAR(50) NOT NULL,
    description TEXT NOT NULL,
    price FLOAT NOT NULL,
    image TEXT NOT NULL,
    stock_quantity INT NOT NULL,
    tax_percetange INT NOT NULL
);

-- +goose Down
DROP TABLE products;