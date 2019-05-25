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
		Root:   &root,
		Output: "",
	}
	Insert(40, tree.Root)
	Insert(60, tree.Root)
	Insert(45, tree.Root)
	Insert(65, tree.Root)
	Insert(20, tree.Root)

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

	tests := []struct {
		input    int
		expected bool
	}{
		{50, true},
		{20, true},
		{65, true},
		{45, true},
		{100, false},
	}

	for _, test := range tests {
		if found, _, _ := Search(test.input, tree.Root, nil); found != test.expected {
			t.Errorf("Search failed")
		}
	}
}

func TestDelete(t *testing.T) {
	// Case 1: delete a node that has no children (here just assign it to nil)
	tree := Setup()
	Insert(15, tree.Root)

	deleted := Delete(15, tree.Root)
	newSubRoot := tree.Root.Left.Left.Left
	if deleted != true && newSubRoot != nil {
		t.Errorf("Case 1: Deletion of node 15 failed")
	}
	if found, _, _ := Search(15, tree.Root, nil); found != false {
		t.Errorf("Case 1: Deletion of node 15 failed")
	}

	// Case 2a: delete a node with a R child, and no L child (here R child should become new sub-root)
	tree = Setup()
	Insert(15, tree.Root)
	Insert(18, tree.Root)

	deleted = Delete(15, tree.Root)
	newSubRoot = tree.Root.Left.Left.Left
	if deleted != true && newSubRoot.Data != 18 {
		t.Errorf("Case 2a: Deletion of node 15 failed")
	}
	if newSubRoot.Right != nil && newSubRoot.Left != nil {
		t.Errorf("Case 2a: Deletion of node 15 failed")
	}
	if found, _, _ := Search(15, tree.Root, nil); found != false {
		t.Errorf("Case 2a: Deletion of node 15 failed")
	}

	// Case 2b: delete a node with a R child and L child (here R child should become new sub-root)
	tree = Setup()
	Insert(15, tree.Root)
	Insert(18, tree.Root)
	Insert(10, tree.Root)

	deleted = Delete(15, tree.Root)
	newSubRoot = tree.Root.Left.Left.Left
	if deleted != true && newSubRoot.Data != 18 {
		t.Errorf("Case 2b: Deletion of node 15 failed")
	}
	if newSubRoot.Right != nil && newSubRoot.Left.Data != 10 {
		t.Errorf("Case 2b: Deletion of node 15 failed")
	}
	if found, _, _ := Search(15, tree.Root, nil); found != false {
		t.Errorf("Case 2b: Deletion of node 15 failed")
	}
	if found, _, _ := Search(18, tree.Root, nil); found != true {
		t.Errorf("Case 2b: Deletion of node 15 failed")
	}
	if found, _, _ := Search(10, tree.Root, nil); found != true {
		t.Errorf("Case 2b: Deletion of node 15 failed")
	}

	// Case 3: delete a node with only a L child (here L child should become new sub-root)
	tree = Setup()
	Insert(15, tree.Root)
	Insert(10, tree.Root)

	deleted = Delete(15, tree.Root)
	newSubRoot = tree.Root.Left.Left.Left
	if deleted != true && newSubRoot.Data != 10 {
		t.Errorf("Case 3: Deletion of node 15 failed")
	}
	if newSubRoot.Right != nil && newSubRoot.Left != nil {
		t.Errorf("Case 3: Deletion of node 15 failed")
	}
	if found, _, _ := Search(15, tree.Root, nil); found != false {
		t.Errorf("Case 3: Deletion of node 15 failed")
	}
}
