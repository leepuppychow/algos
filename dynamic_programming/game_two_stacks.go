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
	Left  *Node // note for a recursive value like this, you need to use a pointer
	Right *Node
	Val   int
	Sum   int
	Depth int
}

var root = Node{
	Left:  nil,
	Right: nil,
	Val:   0,
	Sum:   0,
	Depth: 0,
}

func SolveTwoStacks(maxSum int, leftStack, rightStack []int) int {
	path := ""
	path, _, _, _, _, _, _ = DecisionTree(path, maxSum, leftStack, rightStack, 0, 0, &root)
	fmt.Println(path)
	return len(path)
}

func DecisionTree(path string, maxSum int, leftStack, rightStack []int, indexL, indexR int, current *Node) (string, int, []int, []int, int, int, *Node) {

	fmt.Println(current)
	fmt.Println("INDEX R", indexR)
	fmt.Println("INDEX L", indexL)
	bufio.NewReader(os.Stdin).ReadBytes('\n')

	if current.Left == nil && indexL < len(leftStack) && current.Sum+leftStack[indexL] <= maxSum {
		current.Left = &Node{
			Left:  nil,
			Right: nil,
			Val:   leftStack[indexL],
			Sum:   current.Sum + leftStack[indexL],
			Depth: current.Depth + 1,
		}
		path += strconv.Itoa(leftStack[indexL])
		indexL++
		return DecisionTree(path, maxSum, leftStack, rightStack, indexL, indexR, current.Left)
	}
	if current.Right == nil && indexR < len(rightStack) && current.Sum+rightStack[indexR] <= maxSum {
		current.Right = &Node{
			Left:  nil,
			Right: nil,
			Val:   rightStack[indexR],
			Sum:   current.Sum + rightStack[indexR],
			Depth: current.Depth + 1,
		}
		path += strconv.Itoa(rightStack[indexR])
		indexR++
		return DecisionTree(path, maxSum, leftStack, rightStack, indexL, indexR, current.Right)
	}

	indexL = 0
	indexR = 0
	return DecisionTree(path, maxSum, leftStack, rightStack, indexL, indexR, &root)
}
