package main

import "fmt"

type LRUCache struct {
	Cache    map[int]int
	Queue    []int  // Should use a different data structure (inserting into middle of array is getting costly)
	Capacity int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		Cache:    make(map[int]int),
		Queue:    []int{},
		Capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	fmt.Println("GET", key)

	if value, ok := this.Cache[key]; ok {
		var indexInQueue int
		for i, k := range this.Queue {
			if k == key {
				indexInQueue = i
				break
			}
		}
		if indexInQueue != 0 {
			firstSection := append([]int{key}, this.Queue[:indexInQueue]...)
			secondSection := this.Queue[indexInQueue+1 : len(this.Queue)]
			this.Queue = append(firstSection, secondSection...)
		}
		fmt.Println(this)
		return value
	} else {
		return -1
	}
}

func (this *LRUCache) Set(key, value int) {
	fmt.Println("SET", key, value)
	_, dontDelete := this.Cache[key]
	this.Cache[key] = value
	this.Queue = append([]int{key}, this.Queue...) // this will prepend to the queue
	if len(this.Queue) > this.Capacity {
		keyToDelete := this.Queue[this.Capacity]
		if !dontDelete {
			delete(this.Cache, keyToDelete)
		}
		this.Queue = this.Queue[:this.Capacity]
	}
	fmt.Println(this)
}

func main() {
	lru := Constructor(10)
	lru.Set(7, 28)
	lru.Set(7, 1)
	lru.Set(8, 15)
	lru.Get(6)
	lru.Set(10, 27)
	lru.Set(8, 10)
	lru.Get(8)
	lru.Set(6, 29)
	lru.Set(1, 9)
	lru.Get(6)
	lru.Set(10, 7)
	lru.Get(1)
	lru.Get(2)
	lru.Get(13)
	lru.Set(8, 30)
	lru.Set(1, 5)
	lru.Get(1)
	lru.Set(13, 2)
	lru.Get(12)

	// lru := Constructor(2)
	// lru.Set(2, 1)
	// lru.Set(1, 1)
	// lru.Set(2, 3)
	// lru.Set(4, 1)
	// lru.Get(1)
	// lru.Get(2)
}
