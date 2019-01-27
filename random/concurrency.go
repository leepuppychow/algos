package concurrency

import (
	"github.com/leepuppychow/algos/numbers"
)

func SelectExample() <-chan int {
	ch := make(chan int)
	go func() {
		for {
			select {
			case s := <-numbers.FibonacciFast2(10):
				ch <- s

			}
		}
	}()
	return ch
}
