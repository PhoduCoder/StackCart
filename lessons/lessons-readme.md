## Steps

go mod init lessons 
//Created a module called lessons

This created the go.mod inside the lessons folder

//To get new packages - using Go upstream repository 
go get github.com/jackc/pgx/v5


# Export DB_URL 
export DATABASE_URL='postgres://stackcart:stackcart_dev@127.0.0.1:5432/stackcart?sslmode=disable'

# Query Choices using pgx
QueryRow(...).Scan(...) -> Read one row
Query() -> Read many rows

Exec() -> Execute commands without returning 
