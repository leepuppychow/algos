package backtrack

import (
	"fmt"
	"testing"
)

func TestGetPossibleMoves(t *testing.T) {
	moves := GetPossibleMoves(5, []int{2, 2})
	expected := [][]int{
		[]int{3, 4},
		[]int{4, 3},
		[]int{4, 1},
		[]int{3, 0},
		[]int{1, 0},
		[]int{0, 1},
		[]int{0, 3},
		[]int{1, 4},
	}
	for i, m := range moves {
		expectedMove := expected[i]
		if m[0] != expectedMove[0] || m[1] != expectedMove[1] {
			t.Errorf("GetPossibleMoves failed")
			break
		}
	}
}

func TestGetPossibleMoves3x3(t *testing.T) {
	moves := GetPossibleMoves(3, []int{0, 0})
	expected := [][]int{
		[]int{1, 2},
		[]int{2, 1},
	}

	for i, m := range moves {
		expectedMove := expected[i]
		if m[0] != expectedMove[0] || m[1] != expectedMove[1] {
			t.Errorf("GetPossibleMoves failed")
			break
		}
	}

	moves = GetPossibleMoves(3, []int{2, 1})
	expected = [][]int{
		[]int{0, 0},
		[]int{0, 2},
	}

	for i, m := range moves {
		expectedMove := expected[i]
		if m[0] != expectedMove[0] || m[1] != expectedMove[1] {
			t.Errorf("GetPossibleMoves failed")
			break
		}
	}

	moves = GetPossibleMoves(3, []int{1, 0})
	expected = [][]int{
		[]int{2, 2},
		[]int{0, 2},
	}

	fmt.Println(moves)
	for i, m := range moves {
		expectedMove := expected[i]
		if m[0] != expectedMove[0] || m[1] != expectedMove[1] {
			t.Errorf("GetPossibleMoves failed")
			break
		}
	}

	moves = GetPossibleMoves(3, []int{1, 1})
	expected = [][]int{}
	if len(moves) != 0 {
		t.Errorf("GetPossibleMoves failed")
	}
}
