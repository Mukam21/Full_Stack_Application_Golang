CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,
    type VARCHAR(5) NOT NULL CHECK (char_length(type) >= 2 AND char_length(type) <= 5),
    ticker VARCHAR(30) NOT NULL CHECK (char_length(ticker) >= 1 AND char_length(ticker) <= 30),
    volume REAL NOT NULL,
    price REAL NOT NULL,
    date DATE NOT NULL
);