-- +goose Up
-- +goose StatementBegin
CREATE TABLE
    IF NOT EXISTS Invoice (
        id SERIAL PRIMARY KEY,
        total INTEGER,
        date BIGINT NOT NULL,
        payment_type_id INTEGER NOT NULL REFERENCES Payment_Type(id)
        -- user_id INTEGER REFERENCES Users(id)
    );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS Invoice;
-- +goose StatementEnd