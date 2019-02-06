package main

import (
	"fmt"
	"time"

	dp "github.com/leepuppychow/algos/dynamic_programming"
)

func timeTestTwoVars(fn func(a, b int) int, a, b int) {
	start := time.Now()
	fmt.Println(fn(a, b))
	fmt.Println(time.Since(start))
}

func timeTestOneVar(fn func(a int) int, a int) {
	start := time.Now()
	fmt.Println(fn(a))
	fmt.Println(time.Since(start))
}

func main() {
	// backtrack.RunKnightsTour(3, 0, 0)
	maxSum := 10
	leftStack := []int{4, 2, 4, 6, 1}
	rightStack := []int{2, 1, 8, 5}

	// maxSum = 10
	// leftStack = []int{1, 6, 1, 1, 1}
	// rightStack = []int{2, 1, 7, 4}
	dp.SolveTwoStacks(maxSum, leftStack, rightStack)
}
