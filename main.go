package main

import (
	"fmt"
	"time"

	"github.com/leepuppychow/algos/numbers"
)

func main() {
	start := time.Now()
	fmt.Println(numbers.FibonacciSlow(42))
	fmt.Println(time.Since(start))
}
