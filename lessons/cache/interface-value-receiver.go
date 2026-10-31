package main

import "fmt"

type CacheClient interface {
	Get(key string) string
	Set(key string, value string)
}

type MemoryCache struct {
	data map[string]string
}

func (m MemoryCache) Get(key string) string {
	return m.data[key]
}

func (m MemoryCache) Set(key string, value string) {
	m.data[key] = value
}

func NewCache() MemoryCache {
	return MemoryCache{
		data: make(map[string]string),
	}
}

func main() {
	memory := NewCache()

	fmt.Println(memory.Get("app"))  // Line A
	var client CacheClient = memory // Line B doesn't compile
	// because MemoryCache does not implement CacheClient (Get method has pointer receiver)

	client.Set("app", "Gaurav")

	fmt.Println(client.Get("app"))

	//var client CacheClient = &memory // Line C compiles because *MemoryCache implements CacheClient

}
