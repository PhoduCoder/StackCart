package main

import "fmt"

// Remember this distinction
// - Calling a method: Go can automatically take the address.
// - Satisfying an interface: Go checks the method set of the exact type being assigned.

type CacheClient interface {
	Get(key string) string
	Set(key string, value string)
}

type Cache struct {
	data map[string]string
}

func (m *Cache) Get(key string) string {
	return m.data[key]
}

func (m *Cache) Set(key string, value string) {
	m.data[key] = value
}

func NewCache() *Cache {
	return &Cache{
		data: make(map[string]string),
	}
}

func main() {
	//returns a pointer to Cache.
	memory := NewCache()

	// var client CacheClient = memory
	// This will cause a compile-time error because MemoryCache does not implement CacheClient

	// To fix this, we need to use a pointer to MemoryCache
	var client CacheClient = memory // This works because *MemoryCache implements CacheClient

	client.Set("app", "Grafana")

	fmt.Println(client.Get("app"))
}

/*
Method receiver and interface compatibility
Method declared on	Cache satisfies interface?	*Cache satisfies interface?
(c Cache)			Yes							Yes
(c *Cache)			No							Yes
*/
