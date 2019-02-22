package cache

import (
	"fmt"
	"strconv"
)

// have the cache store Nodes, that way lookup is O(1), do not need to traverse linked list
type LRUCache struct {
	Cache    map[int]*Node
	Capacity int
	Head     *Node
	Tail     *Node
}

type Node struct {
	Key      int
	Value    int
	Previous *Node
	Next     *Node
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		Cache:    make(map[int]*Node),
		Head:     nil,
		Tail:     nil,
		Capacity: capacity,
	}
}

func (this *LRUCache) NewHead(node *Node) {
	node.Next = this.Head
	this.Head.Previous = node
	this.Head = node
}

func (this *LRUCache) NewTail() {
	delete(this.Cache, this.Tail.Key)
	this.Tail.Previous.Next = nil
	this.Tail = this.Tail.Previous
}

func ConnectNeighbors(node *Node) {
	node.Previous.Next = node.Next
	node.Next.Previous = node.Previous
}

func (this *LRUCache) Set(key, value int) {
	fmt.Println("SET", key, value)

	_, inCache := this.Cache[key]
	cacheSize := len(this.Cache)

	if !inCache {
		newNode := Node{
			Key:      key,
			Value:    value,
			Previous: nil,
			Next:     nil,
		}
		this.Cache[key] = &newNode
		if cacheSize == 0 {
			this.Head = &newNode
			this.Tail = &newNode
		} else if cacheSize < this.Capacity {
			this.NewHead(&newNode)
		} else if cacheSize == this.Capacity {
			this.NewHead(&newNode)
			this.NewTail()
		}
	} else {
		existingNode := this.Cache[key]
		existingNode.Value = value
		if existingNode == this.Head {
			// do nothing, update value and leave linked list order the same
		} else if existingNode == this.Tail {
			this.NewHead(existingNode)
			this.NewTail()
		} else { // We are somewhere in the middle of the linked list
			ConnectNeighbors(existingNode)
			this.NewHead(existingNode)
		}
	}

	this.PrintInfo()
}

func (this *LRUCache) Get(key int) int {
	fmt.Println("GET", key)

	return 1000
}

// Helper functions for debugging purposes
func (this *Node) PrintList(result string) string {
	if this == nil {
		return result
	}
	result += strconv.Itoa(this.Key)
	return this.Next.PrintList(result)
}

func (this *LRUCache) PrintInfo() {
	fmt.Println("CACHE", this.Cache)
	fmt.Println("List head", this.Head)
	fmt.Println("List tail", this.Tail)
	fmt.Println("LRU LIST:", this.Head.PrintList(""))
}
