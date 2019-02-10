package linked_list

import "fmt"

type Node struct {
	Data  int
	Child *Node
}

type LinkedList struct {
	Head *Node
}

func (list *LinkedList) Search(current *Node, number int) bool {
	if current.Data == number {
		return true
	} else if current.Child == nil {
		return false
	}
	return list.Search(current.Child, number)
}

func (list *LinkedList) Length() int {
	length := 1
	for current := list.Head; current.Child != nil; current = current.Child {
		length++
	}
	return length
}

func (list *LinkedList) Append(current *Node, node *Node) {
	if current.Child == nil {
		current.Child = node
		return
	}
	list.Append(current.Child, node)
}

func (list *LinkedList) Delete(current *Node, number int) {
	if current.Child.Data == number {
		temp := current.Child
		current.Child = current.Child.Child
		temp.Child = nil
		temp.Data = 0
		return
	}
	list.Delete(current.Child, number)
}

func (list *LinkedList) Contents() bool {
	var current *Node
	for current = list.Head; current.Child != nil; current = current.Child {
		fmt.Printf("%d, ", current.Data)
	}
	fmt.Println(current.Data)
	return true
}

func ReverseList(head *Node) *Node { // Iterative
	current := head
	var past *Node = nil
	var next *Node = nil

	for current != nil {
		next = current.Child
		current.Child = past
		past = current
		current = next
	}
	return past
}
