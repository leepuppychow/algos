// https://www.geeksforgeeks.org/the-knights-tour-problem-backtracking-1/

// Given n (size of the board) and a start position (x, y)
// Can you visit every square on that "chess" board with a Knight's movements
// You can only visit a square once!

// 8x8 board can be covered starting at 0,0
// 5x5 board can be covered starting at 2,2

/* Algorithm:

~ Break recursion if len(Moves) == n*n (Covered whole board)
~ Break recursion if len(Moves) == 1 (Failed to cover board)

1. Initialize matrix for the board with F for all squares (not visited)
2. Get array of all possible Moves from there (8)
		--> x +/- 1 or x +/- 2
		--> y +/- 1 or y +/- 2
3. Exclude non-possibilities
		A. if x or y is negative (out of bounds)
		B. coord already in Moves try next)
4. Iterate through array of possibilties, push into Moves and call KnightFn(Moves)
5. If there are no possible moves, call KnightFn(Moves.pop(lastMove))

*/

package backtrack

import (
	"bufio"
	"fmt"
	"os"
)

func RunKnightsTour(n, startX, startY int) ([][]int, [][]bool) {
	board := [][]bool{}
	for i := 0; i < n; i++ {
		row := make([]bool, n)
		board = append(board, row)
	}
	board[startX][startY] = true
	firstPosition := []int{startX, startY}
	movesTaken := [][]int{firstPosition}

	_, movesTaken, finalBoard := KnightsTour(n, movesTaken, board)
	return movesTaken, finalBoard
}

func GeneratePossibleMoves(position []int, board [][]bool) [][]int {
	n := len(board)
	x := position[0]
	y := position[1]
	moveOptions := [][]int{}

	if x+1 < n && y+2 < n && !board[x+1][y+2] {
		moveOptions = append(moveOptions, []int{x + 1, y + 2})
	}
	if x+2 < n && y+1 < n && !board[x+2][y+1] {
		moveOptions = append(moveOptions, []int{x + 2, y + 1})
	}
	if x+2 < n && y-1 >= 0 && !board[x+2][y-1] {
		moveOptions = append(moveOptions, []int{x + 2, y - 1})
	}
	if x+1 < n && y-2 >= 0 && !board[x+1][y-2] {
		moveOptions = append(moveOptions, []int{x + 1, y - 2})
	}
	if x-1 >= 0 && y-2 >= 0 && !board[x-1][y-2] {
		moveOptions = append(moveOptions, []int{x - 1, y - 2})
	}
	if x-2 >= 0 && y-1 >= 0 && !board[x-2][y-1] {
		moveOptions = append(moveOptions, []int{x - 2, y - 1})
	}
	if x-2 >= 0 && y+1 < n && !board[x-2][y+1] {
		moveOptions = append(moveOptions, []int{x - 2, y + 1})
	}
	if x-1 >= 0 && y+2 < n && !board[x-1][y+2] {
		moveOptions = append(moveOptions, []int{x - 1, y + 2})
	}
	return moveOptions
}

func KnightsTour(n int, movesTaken [][]int, board [][]bool) (int, [][]int, [][]bool) {
	if len(movesTaken) == n*n {
		fmt.Println("********* SOLUTION FOUND **********")
		return n, movesTaken, board
	}
	if len(movesTaken) == 0 {
		fmt.Println("********* NO SOLUTION FOUND **********")
		return n, movesTaken, board
	}

	currentPosition := movesTaken[len(movesTaken)-1]
	x := currentPosition[0]
	y := currentPosition[1]
	possibleMoves := GeneratePossibleMoves(currentPosition, board)

	bufio.NewReader(os.Stdin).ReadBytes('\n')
	fmt.Println("movesTaken", movesTaken)
	fmt.Println("possibleMoves", possibleMoves)

	if len(possibleMoves) == 0 { // backtrack
		board[x][y] = false
		return KnightsTour(n, movesTaken[:len(movesTaken)-1], board)
	} else {
		nextPosition := possibleMoves[0]
		x = nextPosition[0]
		y = nextPosition[1]
		board[x][y] = true
		movesTaken = append(movesTaken, nextPosition)
		return KnightsTour(n, movesTaken, board)
	}
}
