package main

import (
	"fmt"
	"time"

	"github.com/leepuppychow/algos/numbers"
	"github.com/leepuppychow/algos/random"
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
	for num := 1000; num < 1200; num += 3 {
		<-random.ConcurrencySelectExample(num)
	}

	timeTestTwoVars(numbers.GCD, 3918848, 1653264)
	timeTestOneVar(numbers.FibonacciSlow, 43)
	timeTestOneVar(numbers.FibonacciFast, 43)
}
