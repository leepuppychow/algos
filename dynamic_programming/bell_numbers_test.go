package dp

import (
	"testing"
)

func TestBellNumber(t *testing.T) {
	initialRow := []int{1}

	res, _ := BellNumber(1, initialRow)
	if res != 1 {
		t.Errorf("BellNumber failed")
	}
	res, _ = BellNumber(2, initialRow)
	if res != 2 {
		t.Errorf("BellNumber failed")
	}
	res, _ = BellNumber(5, initialRow)
	if res != 52 {
		t.Errorf("BellNumber failed")
	}
	res, _ = BellNumber(12, initialRow)
	if res != 4213597 {
		t.Errorf("BellNumber failed")
	}
}
