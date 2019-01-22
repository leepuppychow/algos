package arrays

import (
	"testing"
)

func TestMaxPairProduct(t *testing.T) {
	arr1 := []int{1, 2, 5, 7, 10, 3, 4, 7}
	if MaxPairProduct(arr1) != 70 {
		t.Errorf("MaxPairProduct test failed")
	}
}
