package cache

import (
	"fmt"
	"strconv"
)

// have the cache store Nodes, that way lookup is O(1), do not need to traverse linked list
type LRUCache struct {
	Cache    map[int]Node
	LruList  *Node
	Capacity int
}

// Use a doubly linked list here, maybe store the Tail too so delete can be faster?
// type LRUList struct {
// 	Head *Node
// 	Tail *Node
// }

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
	currentNode, alreadyThere := this.Cache[key]

	if alreadyThere {
		currentNode.Value = value
	} else {
		currentNode = Node{
			Key:      key,
			Value:    value,
			Previous: nil,
			Next:     nil,
		}
	}
	this.Cache[key] = currentNode

	// New key/value pair becomes head of doubly linked list
	currentNode.Next = this.LruList
	this.LruList.Previous = &currentNode
	this.LruList = &currentNode

	if alreadyThere || len(this.Cache) > this.Capacity {
		this.LruList.DeleteTail()
	}

	fmt.Println(this)
	fmt.Println("LRU LIST:", this.LruList.PrintLruList(""))
}

func (this *Node) PrintLruList(result string) string {
	if this.Next == nil {
		return result
	}
	result += strconv.Itoa(this.Key)
	return this.Next.PrintLruList(result)
}

func (this *Node) DeleteTail() { // Find way to improve this --> O(n)
	if this.Next == nil {
		this.Previous.Next = nil
		return
	}
	this.Next.DeleteTail()
}
