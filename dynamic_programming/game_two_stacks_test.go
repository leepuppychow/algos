package dp

import (
	"testing"
)

func TestSolveTwoStacks(t *testing.T) {
	maxSum := 10
	leftStack := []int{4, 2, 4, 6, 1}
	rightStack := []int{2, 1, 8, 5}

	if SolveTwoStacks(maxSum, leftStack, rightStack) != 4 {
		t.Errorf("Solve Two Stacks failed")
	}
}

func TestSolveTwoStacks2(t *testing.T) {
	maxSum := 10
	leftStack := []int{1, 6, 1, 1, 1}
	rightStack := []int{2, 1, 7, 4}

	if SolveTwoStacks(maxSum, leftStack, rightStack) != 5 {
		t.Errorf("Solve Two Stacks failed")
	}
}
