package numbers

import (
	"testing"
)

func TestGCDEuclideanAlgo(t *testing.T) {
	if GCD(20, 10) != 10 {
		t.Errorf("GCD failed")
	}
	if GCD(3918848, 1653264) != 61232 {
		t.Errorf("GCD failed")
	}
}

func TestLargestDivisor(t *testing.T) {
	if <-LargestDivisor(10) != 5 {
		t.Errorf("Largest Divisor failed")
	}
	if <-LargestDivisor(1233) != 411 {
		t.Errorf("Largest Divisor failed")
	}
	if <-LargestDivisor(3) != 3 {
		t.Errorf("Largest Divisor failed")
	}
	if <-LargestDivisor(12345) != 4115 {
		t.Errorf("Largest Divisor failed")
	}
}
