package main

import (
	"fmt"
	"time"

	"github.com/leepuppychow/algos/numbers"
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
	timeTestTwoVars(numbers.GCD, 3918848, 1653264)
	timeTestOneVar(numbers.FibonacciSlow, 43)
	timeTestOneVar(numbers.FibonacciFast, 43)
}
