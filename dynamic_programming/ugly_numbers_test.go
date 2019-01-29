package dp

import (
	"testing"
)

func UglyNumberTest(t *testing.T) {
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
