
## First Slice
```
Create product → Retrieve product → List products
```

## First Table 
# Product
| Field            | Purpose                           |
| ---------------- | --------------------------------- |
| `id`             | Internal database identity        |
| `sku`            | Unique business identifier        |
| `name`           | Product name                      |
| `description`    | Optional details                  |
| `price_cents`    | Price stored as integer cents     |
| `stock_quantity` | Available inventory               |
| `is_active`      | Whether customers can purchase it |
| `created_at`     | Creation timestamp                |


```
INSERT INTO table_name (column1, column2)
VALUES (value1, value2)
RETURNING *;
```
