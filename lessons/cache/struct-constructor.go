package main

import "fmt"

type Cache struct {
	data map[string]string
	hits int
	name string
}

func NewCache(name string) *Cache {
	return &Cache{
		data: make(map[string]string),
		hits: 0,
		name: name,
	}
}

func (c *Cache) Get(key string) (string, bool) {
	value, exists := c.data[key]
	if exists {
		c.hits++
	}
	return value, exists
}

func (c *Cache) Set(key string, value string) {
	c.data[key] = value
}

func (c Cache) SetNoPointer(key string, value string) {
	c.data[key] = value
}

func (c Cache) GetNoPointer(key string) (string, bool) {
	value, exists := c.data[key]
	if exists {
		c.hits++
	}
	return value, exists
}

func main() {

	//Now instead of writing
	//cache := Cache{
	//	data: make(map[string]string),
	//	hits: 0,
	//	name: "RediCache",
	//}
	//we can simply write

	cache := NewCache("MyCache")
	// This helps to hide initialization details and
	// provides a clear and consistent way to create instances of the struct.

	cache.Set("app", "ArgoCD")

	fmt.Println(cache.Get("app"))

	cache.SetNoPointer("app", "Grafana")

	cache.GetNoPointer("app") //(&cache).Get("app") should have been used

	fmt.Println(cache.Get("app"))

}

/*
For a method call, whether the method receives a copy of the struct or
 operates on the original depends on its receiver type, not the constructor's return type.

- Value receiver (c Cache) → Receives a copy of the struct.
- Pointer receiver (c *Cache) → Receives a pointer to the original struct.

The constructor independently determines whether it returns a Cache value or a *Cache pointer.

All four combinations
Constructor returns	Method receiver	Changes to c.hits persist?
Cache	(c Cache)	No
Cache	(c *Cache)	Yes
*Cache	(c Cache)	No
*Cache	(c *Cache)	Yes

This assumes we assign the constructor result to a local variable before calling the method.

One nuance: A value receiver copies the struct, but fields such as maps can still reference shared underlying data, as we learned earlier.
One final concept to remember
Constructor return type = what the caller receives.

Method receiver type = how the method accesses the struct.

These are separate decisions.

*/

/*
 go run struct-constructor.go
ArgoCD true
Grafana true
*/

//Here you might be confused why line 63 sends a copy of struct to the method,
//  but line 65 still prints Grafana.
// The reason is that the data field is a map, and maps are reference types in Go.
//  When you pass a struct with a map field to a method with a value receiver,
//  the method receives a copy of the struct, but both the original and the copy share the same underlying map.
//  Therefore, changes made to the map in the method affect the original struct's map.
