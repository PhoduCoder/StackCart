package main

type CacheClient interface {
	Get(key string) string
}

type MemoryCache struct {
	data map[string]string
}

func (m *MemoryCache) Get(key string) string {
	return m.data[key]
}

func NewCache() MemoryCache {
	return MemoryCache{
		data: make(map[string]string),
	}
}

func main() {
	memory := NewCache()

	memory.Get("app")               // Line A
	var client CacheClient = memory // Line B doesn't compile
	// because MemoryCache does not implement CacheClient (Get method has pointer receiver)

	var client CacheClient = &memory // Line C compiles because *MemoryCache implements CacheClient

}
