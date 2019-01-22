package numbers

import (
	"fmt"
	"testing"
)

func TestFibonacciSlow(t *testing.T) {
	if FibonacciSlow(3) != 2 {
		t.Errorf("Fibonacci slow failed")
	}
}
