package greedy

import "testing"

func TestEgyptianFraction(t *testing.T) {
	testFrac := Fraction{
		N: 2,
		D: 3,
	}
	expected := []Fraction{
		Fraction{
			N: 1,
			D: 2,
		},
		Fraction{
			N: 1,
			D: 6,
		},
	}

	_, results := EgyptianFraction(testFrac, []Fraction{})
	for i, _ := range results {
		if results[i] != expected[i] {
			t.Errorf("EGYPTIAN FRACTION FAILED")
		}
	}
}

func TestEgyptianFraction2(t *testing.T) {
	testFrac := Fraction{
		N: 6,
		D: 14,
	}
	expected := []Fraction{
		Fraction{
			N: 1,
			D: 3,
		},
		Fraction{
			N: 1,
			D: 11,
		},
		Fraction{
			N: 1,
			D: 231,
		},
	}

	_, results := EgyptianFraction(testFrac, []Fraction{})
	for i, _ := range results {
		if results[i] != expected[i] {
			t.Errorf("EGYPTIAN FRACTION FAILED")
		}
	}
}
