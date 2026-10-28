package main

import "fmt"

func main() {

	var cache map[string]string

	cache1 := make(map[string]string)

	fmt.Println("From make one")
	fmt.Println(len(cache1))
	fmt.Println(cache1["app"] == "")
	fmt.Println(cache1 == nil)
	fmt.Println("This declares a map and initialized with zero values, so in case of string it is \"\" ")

	fmt.Println("===============")
	fmt.Println("From var initialized map")
	fmt.Println(len(cache))
	fmt.Println(cache["app"] == "")
	fmt.Println(cache == nil)
	fmt.Println("This declares a nil map")
}
