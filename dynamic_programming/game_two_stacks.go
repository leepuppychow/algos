package dp

// https://www.hackerrank.com/challenges/game-of-two-stacks/problem

// Problem: Given 2 stacks and a max sum,
// what are the max number of items you can take without exceeding that max value.

/* EXAMPLE: (max sum is 10)

	4
	2				2
	4				1
	6				8
	1				5
------  ------
	A				B

	Possible solutions:
		4, 2, 2, 1 --> 4 moves
		2, 1, 4, 2 --> 4 moves

	ALGORITHM? Try to construct a binary decision tree.
	** Greedy approach taking the min of top of each stack won't work here!
	Each node stores:
		* Stack A child node
		* Stack B child node
		* Value of node
		* CurrentSum at node
		* Depth (this would be equivalent to the score at current position)

	Then traverse tree looking for the max depth without exceeding that sum
*/

type Node struct {
	ChildA *Node // note for a recursive value like this, you need to use a pointer
	ChildB *Node
	Val    int
	Sum    int
	Depth  int
}

func SolveTwoStacks(stackA, stackB []int, max int) int {
	tree := MakeDecisionTree(stackA, stackB)
	depth := TraverseForMaxDepth(tree)
	return depth
}

func MakeDecisionTree(stackA, stackB []int) Node { // return the root of the tree
	root := Node{
		ChildA: nil,
		ChildB: nil,
		Val:    0,
		Sum:    0,
		Depth:  0,
	}
	// current := root

	return root
}

func TraverseForMaxDepth(root Node) int { // Tree traversal - Time O(n)
	return 5
}
