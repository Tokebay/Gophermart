-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS Users (
    id SERIAL PRIMARY KEY,
    login VARCHAR(255) UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    -- Дополнительные поля пользователя, если нужно
);


CREATE TABLE IF NOT EXISTS Transactions (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES Users(id),
    amount INTEGER,
    transaction_type VARCHAR(10), -- "credit" для начисления, "debit" для списания
    timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
