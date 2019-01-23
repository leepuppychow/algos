package numbers

import (
	"testing"
)

func TestFibonacciSlow(t *testing.T) {
	if FibonacciSlow(0) != 0 {
		t.Errorf("Fibonacci slow failed")
	}
	if FibonacciSlow(1) != 1 {
		t.Errorf("Fibonacci slow failed")
	}
	if FibonacciSlow(3) != 2 {
		t.Errorf("Fibonacci slow failed")
	}
	if FibonacciSlow(12) != 144 {
		t.Errorf("Fibonacci slow failed")
	}
}

func TestFibonacciFast(t *testing.T) {
	if <-FibonacciFast(0) != 0 {
		t.Errorf("Fibonacci slow failed")
	}
	if <-FibonacciFast(1) != 1 {
		t.Errorf("Fibonacci slow failed")
	}
	if <-FibonacciFast(3) != 2 {
		t.Errorf("Fibonacci slow failed")
	}
	if <-FibonacciFast(12) != 144 {
		t.Errorf("Fibonacci slow failed")
	}
}