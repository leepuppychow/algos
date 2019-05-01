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
	return BST{
		Root: &root,
	}
}

func TestInsert(t *testing.T) {
	tree := Setup()
	Insert(40, tree.Root)
	Insert(60, tree.Root)
	Insert(45, tree.Root)
	Insert(65, tree.Root)
	Insert(20, tree.Root)

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
