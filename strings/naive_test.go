package strings

import (
	"fmt"
	"testing"
)

func TestIncludesNaive(t *testing.T) {
	text := "Hello jello"
	pattern := "ello"
	expected := []int{1, 7}
	actual := IncludesNaive(text, pattern)

	for i, _ := range expected {
		if expected[i] != actual[i] {
			t.Errorf("IncludesNaive failed")
		}
	}

	pattern = "l"
	expected = []int{2, 3, 8, 9}
	actual = IncludesNaive(text, pattern)

	for i, _ := range expected {
		if expected[i] != actual[i] {
			t.Errorf("IncludesNaive failed")
		}
	}

	pattern = ""
	expected = []int{}
	actual = IncludesNaive(text, pattern)

	for i, _ := range expected {
		if expected[i] != actual[i] {
			t.Errorf("IncludesNaive failed")
		}
	}

	pattern = " "
	expected = []int{5}
	actual = IncludesNaive(text, pattern)

	for i, _ := range expected {
		if expected[i] != actual[i] {
			t.Errorf("IncludesNaive failed")
		}
	}
	fmt.Println("Expected", expected)
	fmt.Println("Actual", actual)
}
