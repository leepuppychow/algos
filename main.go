package main

import (
	"fmt"
	"time"

	"github.com/leepuppychow/algos/numbers"
)

func main() {
	start := time.Now()
	fmt.Println(<-numbers.FibonacciFast(43))
	fmt.Println(time.Since(start))
	
	start = time.Now()
	fmt.Println(numbers.FibonacciSlow(43))
	fmt.Println(time.Since(start))
}
