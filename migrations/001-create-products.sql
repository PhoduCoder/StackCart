CREATE TABLE  IF NOT EXISTS Products(
  id bigserial not null PRIMARY KEY,
  sku  varchar(255) Unique not null,
  name VARCHAR(255) NOT NULL CHECK (name <> ''),
  description TEXT,
  price_cents INT  NOT NULL CHECK (price_cents >= 0),
  stock_quantity INT  NOT NULL CHECK (stock_quantity >= 0),
  is_active BOOLEAN DEFAULT TRUE NOT NULL,
  created_at timestamp DEFAULT current_timestamp NOT NULL
);

INSERT INTO Products (sku, name, description, price_cents, stock_quantity) VALUES
('SKU001', 'Product 1', 'Description for Product 1', 1000, 50),
('SKU002', 'Product 2', 'Description for Product 2', 2000, 30),
('SKU003', 'Product 3', 'Description for Product 3', 1500, 20);

INSERT INTO Products (sku, name, description, price_cents, stock_quantity) VALUES
('SKU005', 'Product 5', 'Description for Product 5', 1000, 50) returning *;

INSERT INTO Products (sku, name, description, price_cents, stock_quantity) VALUES
('SKU006', 'Product 6', 'Description for Product 6', 1000, 50) returning sku,name,price_cents;

