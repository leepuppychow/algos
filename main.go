package main

import (
	"github.com/leepuppychow/algos/dynamic_programming"
	"fmt"
	"time"
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

	maxSum := 5
	leftStack := []int{4, 11, 16, 0, 18, 17, 9, 13, 7, 12, 16, 19, 2, 15, 5, 13, 1, 10, 0, 8, 0, 6, 16, 12, 15, 7, 1, 6, 19, 16, 2}
	rightStack := []int{5, 8, 11, 16, 6, 0, 5, 11, 7, 9, 8, 6, 3, 3, 4, 8, 17, 14, 9, 5, 15, 15, 1, 19, 10, 0, 12, 8, 11, 9, 11, 18, 17, 14}

	fmt.Println(dp.SolveTwoStacks(maxSum, leftStack, rightStack))

}
