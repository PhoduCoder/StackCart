package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func main() {

	//Creates a connection
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	fmt.Println("==============Using Query Row method==============")
	var name string
	var priceCents int
	row := conn.QueryRow(context.Background(), "SELECT name, price_cents FROM products WHERE sku = $1",
		"SKU001")
	err = row.Scan(&name, &priceCents)

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Product: %s, Price: %d cents\n", name, priceCents)

	fmt.Println("==============Using Query method==============")

	rows, err := conn.Query(context.Background(), "select sku, id, name, price_cents from Products where price_cents>=$1 ORDER by id ASC", 1500)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Query failed: %v\n", err)
		os.Exit(1)
	}
	defer rows.Close()

	for rows.Next() {
		var sku string
		var id int
		var name string
		var priceCents int
		err = rows.Scan(&sku, &id, &name, &priceCents)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Row scan failed: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Product: %s, ID: %d, Name: %s, Price: %d cents\n", sku, id, name, priceCents)
	}

	//Your loop runs while rows.Next() returns true
	//When rows.Next() returns false, you have finished reading all the rows or
	//an error occurred.
	//You can check for errors after the loop by calling rows.Err().

	//Example say we had a query that returned 10 rows, but the 5th row had an error. The first four rows would be read successfully,
	// and rows.Next() would return true for each of those four rows. When it got to the 5th row, rows.Next() would return false, and the loop would exit. You could then check rows.Err() to see that there was an error reading the 5th row.

	// Check for errors from iterating over rows.
	//rows.Err() == nil: Reading finished normally.
	//rows.Err() != nil: Reading failed; this gives you the error.

	if err := rows.Err(); err != nil {
		fmt.Printf("Error reading results: %v\n", err)
		return
	}
}
