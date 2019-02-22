package cache

import (
	"fmt"
	"strconv"
)

// Have the cache store Nodes, that way lookup is O(1), do not need to traverse linked list
// Use a doubly linked list for tracking the LRU items - O(1) to insert if we know where node is
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

func ConnectNeighbors(node *Node) {
	node.Previous.Next = node.Next
	node.Next.Previous = node.Previous
}

func (this *LRUCache) NewHead(node *Node) {
	node.Next = this.Head
	this.Head.Previous = node
	this.Head = node
}

func (this *LRUCache) NewTail() {
	this.Tail.Previous.Next = nil
	this.Tail = this.Tail.Previous
}

func (this *LRUCache) ReorderList(existingNode *Node) {
	if existingNode == this.Head {
		return // no need to re-arrange if the head is selected
	} else if existingNode == this.Tail {
		this.NewHead(existingNode)
		this.NewTail()
	} else { // We are somewhere in the middle of the linked list
		ConnectNeighbors(existingNode)
		this.NewHead(existingNode)
	}
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
			delete(this.Cache, this.Tail.Key)
			this.NewTail()
		}
	} else {
		existingNode := this.Cache[key]
		existingNode.Value = value
		this.ReorderList(existingNode)
	}

	this.PrintInfo()
}

func (this *LRUCache) Get(key int) int {
	result := -1
	existingNode, inCache := this.Cache[key]
	if inCache {
		result = existingNode.Value
		this.ReorderList(existingNode)
	}

	fmt.Printf("GET %d; VALUE: %d\n", key, result)
	this.PrintInfo()
	return result
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
	fmt.Println("\tCACHE", this.Cache)
	fmt.Println("\tList head", this.Head)
	fmt.Println("\tList tail", this.Tail)
	fmt.Println("\tLRU LIST:", this.Head.PrintList(""))
}
