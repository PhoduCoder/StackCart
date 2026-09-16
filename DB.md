## Run postgres locally 

```
docker run --name stackcart-postgres \
  -e POSTGRES_USER=stackcart \
  -e POSTGRES_PASSWORD=stackcart_dev \
  -e POSTGRES_DB=stackcart \
  -p 127.0.0.1:5432:5432 \
  -v stackcart-pg18-data:/var/lib/postgresql \
  -d postgres:18
```

## Exec to the container
```
docker exec -it stackcart-postgres psql -U stackcart -d stackcart
```

## Check the schema of a table 
```
\d products 
SELECT * FROM products;

```

\q - Quit out of postgres


```
INSERT INTO Products (sku, name, description, price_cents, stock_quantity) VALUES
('SKU006', 'Product 6', 'Description for Product 5', 1000, 50) returning sku,name,price_cents;
  sku   |   name    | price_cents 
--------+-----------+-------------
 SKU006 | Product 6 |        1000
```

# These are three separate decisions: which columns to return, which rows to include, and how to order them.

```
select sku, id, name, price_cents from Products where price_cents>15 ORDER by id ASC;
```

## Golang with Postgres
https://github.com/jackc/pgx/wiki/Getting-started-with-pgx


