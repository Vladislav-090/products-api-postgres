CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    title VARCHAR(150) NOT NULL,
    price NUMERIC(10, 2) CHECK (price > 0), 
    in_stock BOOLEAN DEFAULT true, 
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

