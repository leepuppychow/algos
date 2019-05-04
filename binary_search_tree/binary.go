package binary_search_tree

import (
	"strconv"
)

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

type BST struct {
	Root   *Node
	Output string
}

func Insert(data int, node *Node) {
	if data < node.Data {
		if node.Left == nil {
			node.Left = &Node{
				Data:  data,
				Left:  nil,
				Right: nil,
			}
			return
		} else {
			Insert(data, node.Left)
		}
	} else {
		if node.Right == nil {
			node.Right = &Node{
				Data:  data,
				Left:  nil,
				Right: nil,
			}
			return
		} else {
			Insert(data, node.Right)
		}
	}
}

func Search(data int, current, parent *Node) (bool, *Node, *Node) {
	if current.Data == data {
		return true, current, parent
	} else if current.Left == nil && current.Right == nil {
		return false, current, parent
	}

	if data < current.Data {
		return Search(data, current.Left, current)
	} else {
		return Search(data, current.Right, current)
	}
}

func Delete(data int, current *Node) bool {
	found, node, parent := Search(data, current, nil)
	if !found {
		return false
	}
	if node.Right != nil {
		temp := node.Right
		temp.Left = node.Left
		ShiftSubTree(temp, parent, false)
	} else if node.Left != nil {
		temp := node.Left
		ShiftSubTree(temp, parent, false)
	} else {
		ShiftSubTree(node, parent, true)
	}
	return true
}

func ShiftSubTree(current, parent *Node, isLeaf bool) {
	if current.Data > parent.Data {
		if isLeaf {
			parent.Right = nil
		} else {
			parent.Right = current
		}
	} else {
		if isLeaf {
			parent.Left = nil
		} else {
			parent.Left = current
		}
	}
}

func (tree *BST) SortedOutputAsc(current *Node) {
	if current.Left != nil {
		tree.SortedOutputAsc(current.Left)
	}
	tree.Output += strconv.Itoa(current.Data) + " "
	if current.Right != nil {
		tree.SortedOutputAsc(current.Right)
	}
}
