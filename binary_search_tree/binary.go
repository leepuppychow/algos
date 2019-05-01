package binary_search_tree

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

type BST struct {
	Root *Node
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

func Search(data int, current *Node) bool {
	if current.Data == data {
		return true
	} else if current.Left == nil && current.Right == nil {
		return false
	}

	if data < current.Data {
		return Search(data, current.Left)
	} else {
		return Search(data, current.Right)
	}
}

func (tree *BST) Delete() {
	
}
