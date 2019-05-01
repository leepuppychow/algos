package cache

import "fmt"

type LRUCacheV1 struct {
	Cache    map[int]int
	Queue    []int // Should use a different data structure (inserting into middle of array is getting costly)
	Capacity int
}

func ConstructorV1(capacity int) LRUCacheV1 {
	return LRUCacheV1{
		Cache:    make(map[int]int),
		Queue:    []int{},
		Capacity: capacity,
	}
}

func (this *LRUCacheV1) Get(key int) int {
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

func (this *LRUCacheV1) Set(key, value int) {
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
