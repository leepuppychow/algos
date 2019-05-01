package binary_search_tree

type Node struct {
	Data int
	Left *Node
	Right *Node
}

type BST struct {
	Root *Node
}

func Insert(data int, node *Node) {
	if data < node.Data {
		if node.Left == nil {
			node.Left = &Node{
				Data: data,
				Left: nil,
				Right: nil,
			}
			return
		} else {
			Insert(data, node.Left)
		}
	} else {
		if node.Right == nil {
			node.Right = &Node{
				Data: data,
				Left: nil,
				Right: nil,
			}
			return
		} else {
			Insert(data, node.Right)
		}
	}
}