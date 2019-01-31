package dp

import (
	"testing"
)

func TestUglyNumber(t *testing.T) {
	if UglyNumber(1) != 1 {
		t.Errorf("UglyNumber function failed")
	}
	if UglyNumber(5) != 5 {
		t.Errorf("UglyNumber function failed")
	}
	if UglyNumber(11) != 15 {
		t.Errorf("UglyNumber function failed")
	}
	if UglyNumber(300) != 82944 {
		t.Errorf("UglyNumber function failed")
	}
}

func TestSlowUglyNumber(t *testing.T) {
	if UglyNumber(1) != 1 {
		t.Errorf("UglyNumber function failed")
	}
	if UglyNumber(5) != 5 {
		t.Errorf("UglyNumber function failed")
	}
	if SlowUglyNumber(11) != 15 {
		t.Errorf("SlowUglyNumber function failed")
	}
	if SlowUglyNumber(300) != 82944 {
		t.Errorf("SlowUglyNumber function failed")
	}
}
