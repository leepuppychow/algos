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
