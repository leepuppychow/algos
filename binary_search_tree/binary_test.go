package binary_search_tree

import (
	"testing"
)

func Setup() BST {
	root := Node{
		Data:  50,
		Left:  nil,
		Right: nil,
	}
	tree := BST{
		Root: &root,
	}
	tree.Insert(40, tree.Root)
	tree.Insert(60, tree.Root)
	tree.Insert(45, tree.Root)
	tree.Insert(65, tree.Root)
	tree.Insert(20, tree.Root)

	return tree
}

func TestInsert(t *testing.T) {
	tree := Setup()

	if tree.Root.Data != 50 {
		t.Errorf("Insert failed, root should be 50, got: %d", tree.Root.Data)
	}
	if tree.Root.Left.Data != 40 {
		t.Errorf("Insert failed")
	}
	if tree.Root.Left.Left.Data != 20 {
		t.Errorf("Insert failed")
	}
	if tree.Root.Left.Right.Data != 45 {
		t.Errorf("Insert failed")
	}
	if tree.Root.Right.Data != 60 {
		t.Errorf("Insert failed")
	}
	if tree.Root.Right.Right.Data != 65 {
		t.Errorf("Insert failed")
	}
	if tree.Root.Right.Left != nil {
		t.Errorf("Insert failed")
	}
}

func TestSearch(t *testing.T) {
	tree := Setup()

	if tree.Search(50, tree.Root) != true {
		t.Errorf("Search failed")
	}
	if tree.Search(20, tree.Root) != true {
		t.Errorf("Search failed")
	}
	if tree.Search(65, tree.Root) != true {
		t.Errorf("Search failed")
	}
	if tree.Search(45, tree.Root) != true {
		t.Errorf("Search failed")
	}
	if tree.Search(100, tree.Root) != false {
		t.Errorf("Search failed")
	}
}

func TestDelete(t *testing.T) {

}
