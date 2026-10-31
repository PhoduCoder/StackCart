package main

import "fmt"

type Cache struct {
	data map[string]string
}

type CacheClient interface {
	Get(key string) (string, bool)
	Set(key string, value string)
}

func (c Cache) Get(key string) (string, bool) {
	value, exists := c.data[key]
	return value, exists
}

func (c Cache) Set(key string, value string) {
	c.data[key] = value
}

func newCache() *Cache {
	return &Cache{
		data: make(map[string]string),
	}
}

func main() {
	cache := newCache()

	cache.Set("app", "ArgoCD")
	cache.Set("version", "3.0")

	fmt.Println(cache.Get("app"))
	fmt.Println(cache.Get("version"))

	var client CacheClient = cache

	fmt.Println(client.Get("app"))
}
