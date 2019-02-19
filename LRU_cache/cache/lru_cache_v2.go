package cache

import (
	"fmt"
	"strconv"
)

type LRUCache struct {
	Cache    map[int]Node
	LruList  *Node // Use a doubly linked list here
	Capacity int
}

type Node struct {
	Key      int
	Value    int
	Previous *Node
	Next     *Node
}

func Constructor(capacity int) LRUCache {
	head := Node{
		Key:      0,
		Value:    0,
		Previous: nil,
		Next:     nil,
	}
	return LRUCache{
		Cache:    make(map[int]Node),
		LruList:  &head,
		Capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	var result int
	if node, ok := this.Cache[key]; ok {
		result = node.Value
	} else {
		result = -1
	}

	fmt.Println("GET", key, result)
	return result
}

func (this *LRUCache) Set(key, value int) {
	fmt.Println("SET", key, value)
	newNode := Node{
		Key:      key,
		Value:    value,
		Previous: nil,
		Next:     nil,
	}
	this.Cache[key] = newNode

	// New key/value pair becomes head of doubly linked list
	newNode.Next = this.LruList
	this.LruList.Previous = &newNode
	this.LruList = &newNode

	fmt.Println(this)
}

func (this *Node) PrintLruList(result string) string {
	if this.Next == nil {
		return result
	}
	result += strconv.Itoa(this.Key)
	return this.Next.PrintLruList(result)
}
