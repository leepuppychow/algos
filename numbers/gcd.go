// Greatest common divisor (using Euclidean algorithm)

package numbers

import (
	"fmt"
	"math"
	"time"
)

func GCD(a, b int) int {
	// Just assume that a is the larger number
	fmt.Println("GCD called")
	if b == 0 {
		return a
	}
	return GCD(b, a%b)
}

func LargestDivisor(a int, ms time.Duration) <-chan int {
	time.Sleep(ms * time.Millisecond) // optional delay for the concurrency example

	ch := make(chan int)
	go func() {
		for i := 2; i <= int(math.Sqrt(float64(a))); i++ {
			if a%i == 0 {
				ch <- a / i
			}
		}
		ch <- a
	}()
	return ch
}
