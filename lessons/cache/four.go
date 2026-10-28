package main

import "fmt"

func change(m map[string]string) {
	delete(m, "app")
	m["database"] = "Redis"
}

func main() {
	cache := make(map[string]string)
	cache["app"] = "ArgoCD"

	change(cache)

	ok, val := cache["app"]
	fmt.Println(ok, val) // prints false and empty string

	fmt.Println(cache["database"]) //redis
}
