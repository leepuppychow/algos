// https://www.geeksforgeeks.org/the-knights-tour-problem-backtracking-1/

// Given n (size of the board) and a start position (x, y)
// Can you visit every square on that "chess" board with a Knight's movements
// You can only visit a square once!

// Sample output:

// (True, [[0,0], [1,2], [3,3] ... ])

/* Algorithm:

~ Break recursion if len(Moves) == n*n (Covered whole board)
~ Break recursion if len(Moves) == 1 (Failed to cover board)

1. initialize the Moves array with start position input
2. Get array of all possible Moves from there (8)
		--> x +/- 1 or x +/- 2
		--> y +/- 1 or y +/- 2
3. Iterate through array of possibilties

	--> if x or y is negative (out of bounds) or coord already in Moves try next
	--> else push into Moves and call KnightFn(Moves)

4. If there are no possible moves, call KnightFn(Moves.pop(lastMove))

*/

package backtrack

import "fmt"

func KnightsTour(n int, moves [][]int) (int, [][]int) {
	if len(moves) == n*n || len(moves) == 0 {
		return n, moves
	}
	currentMove := moves[len(moves)-1]
	possibilities := GetPossibleMoves(n, currentMove)
	for _, p := range possibilities {
		// Now just check if move already includes move
		// if p[0] < 0 || p[0] > maxSize || p[1] < 0 || p[1] > maxSize {
		// 	continue
		// } else {
		// 	moves = append(moves, p)
		// 	return KnightsTour(n, moves)
		// }
		fmt.Println(p)
	}
	return KnightsTour(n, moves[:len(moves)-1])
}

func GetPossibleMoves(size int, move []int) [][]int {
	possibles := [][]int{}
	x := move[0]
	y := move[1]
	// Could refactor into two for loops (-2 to 2)
	// then have if statement in inner loop
	if x+1 < size && y+2 < size {
		possibles = append(possibles, []int{x + 1, y + 2})
	}
	if x+2 < size && y+1 < size {
		possibles = append(possibles, []int{x + 2, y + 1})
	}
	if x+2 < size && y-1 >= 0 {
		possibles = append(possibles, []int{x + 2, y - 1})
	}
	if x+1 < size && y-2 >= 0 {
		possibles = append(possibles, []int{x + 1, y - 2})
	}
	if x-1 >= 0 && y-2 >= 0 {
		possibles = append(possibles, []int{x - 1, y - 2})
	}
	if x-2 >= 0 && y-1 >= 0 {
		possibles = append(possibles, []int{x - 2, y - 1})
	}
	if x-2 >= 0 && y+1 < size {
		possibles = append(possibles, []int{x - 2, y + 1})
	}
	if x-1 >= 0 && y+2 < size {
		possibles = append(possibles, []int{x - 1, y + 2})
	}

	return possibles
}
