package random

import (
	"fmt"

	"github.com/leepuppychow/algos/numbers"
)

func ConcurrencySelectExample(num int) <-chan int {
	ch := make(chan int)
	go func() { // Using channel generator pattern
		for {
			select {
			case s := <-numbers.FibonacciFast2(num):
				fmt.Printf("Fibonacci number (%d) = %d\n", num, s)
				ch <- s
			case s := <-numbers.LargestDivisor(num, 50):
				fmt.Printf("Largest Divisor of (%d) = %d\n", num, s)
				ch <- s
			default:
				fmt.Println("No channel ready yet")
				ch <- 0
			}
		}
	}()
	return ch
}
