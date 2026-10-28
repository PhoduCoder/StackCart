package main

import "fmt"

func main() {
	cache := make(map[string]string)

	cache["app"] = "ArgoCD"
	cache["version"] = "3.0"

	fmt.Println(cache["app"])
	fmt.Println(cache["version"])

	fmt.Println(cache["missing"] == "")

	fmt.Println("===============Two values from map===================")
	val, exists := cache["chart"]
	fmt.Println(val)
	fmt.Println(exists)

	fmt.Println("===============Two values from map===================")
	val1, exists1 := cache["app"] // gives value associated with key and true if key exists otherwise zero value and false
	fmt.Println(val1)
	fmt.Println(exists1)
}
