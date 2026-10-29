package main

import "fmt"

type Cache struct {
	data map[string]string
	hits int
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
	cache := Cache{
		data: make(map[string]string),
		hits: 0,
	}

	cache.Set("app", "ArgoCD") //(&cache).Get("app") should have been used
	//but Go lets us shorten this to cache.Get("app")
	//because the method has a pointer receiver. Go automatically takes the address of cache and passes it to the method.
	cache.Set("version", "3.0")

	cache.SetNoPointer("app", "Grafana")
	cache.SetNoPointer("version", "12.4.0")

	cache.GetNoPointer("app")     //(&cache).Get("app") should have been used
	cache.GetNoPointer("version") //(&cache).Get("version") should have been used
	cache.GetNoPointer("missing") //(&cache).Get("missing") should have been used
	cache.GetNoPointer("app")

	val, exists := cache.Get("app")
	fmt.Println(val, exists)

	val, exists = cache.Get("version")
	fmt.Println(val, exists)

	val, exists = cache.Get("missing")
	fmt.Println(val, exists)

	fmt.Println("Cache hits:", cache.hits)
}

//When you have a pointer receiver, both calls works
//cache.Get("app")     // Go automatically takes the address
//(&cache).Get("app")  // You explicitly take the address

// in this case, even though SetNoPointer and GetNoPointer
// uses value receiver, they are passed a copy of struct
// but the data field is a map, which is a reference type.
// So, when we modify the map through the copy of the struct,
//  we are actually modifying the original map.
// However, the hits field is an int, which is a value type.
// So, when we modify the hits field through the copy of the struct,
// we are modifying a copy of the hits field, not the original hits field.
// This is why the hits field does not get incremented when we call GetNoPointer.
