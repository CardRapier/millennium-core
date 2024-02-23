-- +goose Up
-- +goose StatementBegin
CREATE TABLE
    IF NOT EXISTS Payment_Type (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);
INSERT INTO Payment_Type (name) VALUES ('Cash'), ('Credit Card'), ('Debit Card'), ('Nequi'), ('Daviplata');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS Payment_Type;
-- +goose StatementEnd