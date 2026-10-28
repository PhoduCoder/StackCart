package main

import "fmt"

func getValue(cache map[string]string, key string) string {
	return cache[key]
}

func setValue(cache map[string]string, key string, value string) {
	cache[key] = value
}

func main() {
	cache := make(map[string]string)
	cache["app"] = "ArgoCD"

	setValue(cache, "version", "3.0")

	result := getValue(cache, "app")
	fmt.Println(result)

	val, exists := cache["version"]
	fmt.Println(val)
	fmt.Println(exists)
}
