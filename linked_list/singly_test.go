package linked_list

import (
	"testing"
)

func Setup() LinkedList {
	node3 := Node{
		Data:  25,
		Child: nil,
	}
	node2 := Node{
		Data:  17,
		Child: &node3,
	}
	node1 := Node{
		Data:  3,
		Child: &node2,
	}
	head := Node{
		Data:  5,
		Child: &node1,
	}
	list := LinkedList{Head: &head}
	return list
}

func TestInsert(t *testing.T) {
	list := Setup()
	length := list.Length()
	tests := []int{10, 11, 12, 13, 14, 100, 101}

	for _, test := range tests {
		list.Insert(test)
		if list.Head.Data != test {
			t.Errorf("Insert test failed")
		}
		if list.Length() != length+1 {
			t.Errorf("Insert test failed")
		}
		length++
	}
}

func TestLength(t *testing.T) {
	list := Setup()
	if list.Length() != 4 {
		t.Errorf("Length function failed")
	}
}

func TestSearch(t *testing.T) {
	list := Setup()
	if !list.Search(list.Head, 17) {
		t.Errorf("Search function failed")
	}
	if list.Search(list.Head, 1000) {
		t.Errorf("Search function failed")
	}
}

func TestAppend(t *testing.T) {
	list := Setup()
	newNode := Node{
		Data:  23,
		Child: nil,
	}
	newNode2 := Node{
		Data:  99,
		Child: nil,
	}

	list.Append(list.Head, &newNode)
	list.Append(list.Head, &newNode2)
	if list.Length() != 6 || !list.Search(list.Head, 99) || !list.Search(list.Head, 23) {
		t.Errorf("Append failed")
	}
}

func TestAppendEmptyList(t *testing.T) {
	newHead := Node{
		Data:  0,
		Child: nil,
	}
	newList := LinkedList{Head: &newHead}

	node99 := Node{
		Data:  99,
		Child: nil,
	}
	node98 := Node{
		Data:  98,
		Child: nil,
	}
	newList.Append(newList.Head, &node98)
	newList.Append(newList.Head, &node99)
	if newList.Length() != 3 || !newList.Search(newList.Head, 99) || !newList.Search(newList.Head, 98) {
		t.Errorf("Append failed")
	}
}

func TestDelete(t *testing.T) {
	list := Setup()
	list.Delete(list.Head, 17)
	if list.Length() != 3 || list.Search(list.Head, 17) {
		t.Errorf("Delete failed")
	}
}

func TestContents(t *testing.T) {
	list := Setup()
	if !list.Contents() {
		t.Errorf("Contents failed")
	}
}
