package main

import (
	"fmt"

	c "github.com/leepuppychow/algos/LRU_cache/cache"
)

func main() {
	// lru := c.Constructor(10)
	// lru.Set(7, 28)
	// lru.Set(7, 1)
	// lru.Set(8, 15)
	// lru.Get(6)
	// lru.Set(10, 27)
	// lru.Set(8, 10)
	// lru.Get(8)
	// lru.Set(6, 29)
	// lru.Set(1, 9)
	// lru.Get(6)
	// lru.Set(10, 7)
	// lru.Get(1)
	// lru.Get(2)
	// lru.Get(13)
	// lru.Set(8, 30)
	// lru.Set(1, 5)
	// lru.Get(1)
	// lru.Set(13, 2)
	// lru.Get(12)

	lru := c.Constructor(2)
	lru.Set(2, 1)
	lru.Set(1, 1)
	lru.Set(2, 3)
	lru.Set(4, 1)
	lru.Get(1)
	lru.Get(2)
	lru.Get(200)

	fmt.Println(lru.LruList.PrintLruList(""))
}
