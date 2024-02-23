-- +goose Up
-- +goose StatementBegin
CREATE TABLE Sales (
    id SERIAL PRIMARY KEY,
    product_id INTEGER NOT NULL REFERENCES Product(id),
    invoice_id INTEGER NOT NULL REFERENCES Invoice(id),
    quantity INTEGER NOT NULL,
    price INTEGER NOT NULL,
    iva INTEGER NOT NULL,
    -- gross_margin INTEGER NOT NULL,
    subtotal INTEGER NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS Sales;
-- +goose StatementEnd