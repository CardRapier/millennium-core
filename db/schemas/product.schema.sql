-- +goose Up
-- +goose StatementBegin
CREATE TABLE
    IF NOT EXISTS Product (
        id  SERIAL PRIMARY KEY,
        utc VARCHAR(12) UNIQUE NOT NULL,
        name TEXT NOT NULL,
        price INTEGER NOT NULL,
        gross_margin INTEGER NOT NULL,
        iva INTEGER NOT NULL
    );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS Product;
-- +goose StatementEnd