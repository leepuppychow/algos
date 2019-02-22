package main

import (
	c "github.com/leepuppychow/algos/LRU_cache/cache"
)

func main() {
	lru := c.Constructor(3)
	lru.Set(1, 10)
	lru.Set(2, 20)
	lru.Set(3, 30)
	lru.Set(4, 40)
	lru.Set(5, 50)
	lru.Set(6, 60)
	lru.Set(4, 400)
	lru.Set(6, 600)
	lru.Set(5, 500)
	lru.Set(3, 333)

	lru.Get(4)
	lru.Get(5)
	lru.Get(6)
	lru.Get(3)
}
