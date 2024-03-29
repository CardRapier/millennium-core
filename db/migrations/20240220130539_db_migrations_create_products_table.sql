-- +goose Up
-- +goose StatementBegin
CREATE TABLE
    IF NOT EXISTS Payment_Type (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);
INSERT INTO Payment_Type (name) VALUES ('Cash'), ('Credit Card'), ('Debit Card'), ('Nequi'), ('Daviplata');

CREATE TABLE
    IF NOT EXISTS Invoice (
        id SERIAL PRIMARY KEY,
        total INTEGER,
        date BIGINT NOT NULL,
        payment_type_id INTEGER NOT NULL REFERENCES Payment_Type(id)
        -- user_id INTEGER REFERENCES Users(id)
    );

CREATE TABLE
    IF NOT EXISTS Product (
        id  SERIAL PRIMARY KEY,
        utc VARCHAR(12) UNIQUE NOT NULL,
        name TEXT NOT NULL,
        price INTEGER NOT NULL,
        gross_margin INTEGER NOT NULL,
        iva INTEGER NOT NULL
    );

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
DROP TABLE IF EXISTS Payment_Type;
DROP TABLE IF EXISTS Invoice;
DROP TABLE IF EXISTS Product;
DROP TABLE IF EXISTS Sales;
-- +goose StatementEnd