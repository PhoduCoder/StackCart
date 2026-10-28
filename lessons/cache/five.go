package main

import "fmt"

func main() {

	type Cache struct {
		data map[string]string
	}

	//cache := Cache{}
	//cache.data["app"] = "ArgoCD"
	// The above line will cause a panic because the data field was not initialized
	// This means you have a nil map, and you cannot assign values to a nil map.
	// To fix this, we need to initialize the data field with an empty map using make.
	// You can read from a nil map, but you cannot write to it. Writing to a nil map will cause a runtime panic.

	cache := Cache{
		data: make(map[string]string),
	}

	cache.data["app"] = "ArgoCD"

	fmt.Println(cache.data["app"])
}

// In this example, we define a struct called Cache that contains a field data
//  of type map[string]string. We then create an instance of Cache and initialize the data
//  field with an empty map.
//  We can access and modify the map through the data field of the Cache struct because we used make
//  to create the map. This allows us to encapsulate the map within a struct and provide a more structured way to work with it.
