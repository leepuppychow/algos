package dp

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// https://www.hackerrank.com/challenges/game-of-two-stacks/problem

// Problem: Given 2 stacks and a max sum,
// what are the max number of items you can take without exceeding that max value.

/* EXAMPLE: (max sum is 10)

	4
	2				2
	4				1
	6				8
	1				5
------			  ------
  Left			  Right

	Possible solutions:
		4, 2, 2, 1 --> 4 moves
		2, 1, 4, 2 --> 4 moves

	ALGORITHM? Try to construct a binary decision tree.
	** Greedy approach taking the min of top of each stack won't work here!
	Each node stores:
		* Left Stack child node
		* Right Stack child node
		* Value of node
		* CurrentSum at node
		* Depth (this would be equivalent to the score at current position)

	Then traverse tree looking for the max depth without exceeding that sum
	Try Post Order Traversal (Left, Right, Root)

	1. If there is no left child add new node from top of Left Stack (update its value, currentSum & depth)
		--> increment indexLeft value
		--> continue until currentSum > maxSum or both stacks are empty
		--> if a stack is empty that is ok, just add one child from non-empty stack (L or R)
	2. If there is no right child add new node from top of Right Stack (update its value, currentSum & depth)
		--> increment indexRight value
		--> continue until currentSum > maxSum or both stacks are empty
		--> if a stack is empty that is ok, just add one child from non-empty stack (L or R)
	3. If we have both L and R children then go back to local root and repeat

	In Order = (Left, Root, Right)
	Pre Order = (Root, Left, Right)
*/

type Node struct {
	Left   *Node // note for a recursive value like this, you need to use a pointer
	Right  *Node
	Parent *Node
	Val    int
	Sum    int
	Path   string
	IndexL int
	IndexR int
}

func SolveTwoStacks(maxSum int, leftStack, rightStack []int) int {
	root := Node{
		Left:   nil,
		Right:  nil,
		Parent: nil,
		Val:    0,
		Sum:    0,
		Path:   "",
		IndexL: 0,
		IndexR: 0,
	}
	DecisionTree(maxSum, leftStack, rightStack, &root)
	return len("hello")
}

func DecisionTree(maxSum int, leftStack, rightStack []int, current *Node) (int, []int, []int, *Node) {

	fmt.Println(current)
	bufio.NewReader(os.Stdin).ReadBytes('\n')

	if current.Left == nil && current.IndexL < len(leftStack) && current.Sum+leftStack[current.IndexL] <= maxSum {
		current.Left = &Node{
			Left:   nil,
			Right:  nil,
			Parent: current,
			Val:    leftStack[current.IndexL],
			Sum:    current.Sum + leftStack[current.IndexL],
			Path:   current.Path + strconv.Itoa(leftStack[current.IndexL]),
			IndexL: current.IndexL + 1,
			IndexR: current.IndexR,
		}
		return DecisionTree(maxSum, leftStack, rightStack, current.Left)
	}
	if current.Right == nil && current.IndexR < len(rightStack) && current.Sum+rightStack[current.IndexR] <= maxSum {
		current.Right = &Node{
			Left:   nil,
			Right:  nil,
			Parent: current,
			Val:    rightStack[current.IndexR],
			Sum:    current.Sum + rightStack[current.IndexR],
			Path:   current.Path + strconv.Itoa(rightStack[current.IndexR]),
			IndexR: current.IndexR + 1,
			IndexL: current.IndexL,
		}
		return DecisionTree(maxSum, leftStack, rightStack, current.Right)
	}

	return DecisionTree(maxSum, leftStack, rightStack, current.Parent)
}
