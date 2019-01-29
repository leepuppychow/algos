package greedy

import (
	"testing"
)

func TestActivitySelection(t *testing.T) {
	expectedActivities := []Activity{
		Activity{
			Name:   "A",
			Start:  1,
			Finish: 2,
		},
		Activity{
			Name:   "B",
			Start:  3,
			Finish: 4,
		},
		Activity{
			Name:   "D",
			Start:  5,
			Finish: 7,
		},
		Activity{
			Name:   "E",
			Start:  8,
			Finish: 9,
		},
	}
	num, activities := ActivitySelection(Activities)
	if num != 4 {
		t.Errorf("ActivitySelection Failed")
	}

	for i, a := range activities {
		if a != expectedActivities[i] {
			t.Errorf("ActivitySelection Failed")
		}
	}
}
