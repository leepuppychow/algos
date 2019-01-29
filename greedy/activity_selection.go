// https://www.geeksforgeeks.org/activity-selection-problem-greedy-algo-1/

// Given a set of activities with a start and finish time, you want to pick the maximum activities
// that a person could perform from that set.

// Assumptions:
// 1) Person can only do 1 activity at a time
// 2) Assume activities are already sorted by Finish Times

package greedy

type Activity struct {
	Name   string
	Start  int
	Finish int
}

var Activities = []Activity{
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
		Name:   "C",
		Start:  0,
		Finish: 6,
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
	Activity{
		Name:   "F",
		Start:  5,
		Finish: 9,
	},
}

// Algorithm (greedy approach) --> "picking local optimum will result in global optimum"
// 1. Sort by Finish time (if not already sorted)
// 2. Pick the Activity with lowest finish time (index 0)
// 3. Iterate through array (looking for the lowest finish time),
// 		A. If that activity's start time >= previous activity Finish
// 		B. Then Add to list of selected activities
// 4. Repeat step 3 until we have reached end of array

func ActivitySelection(activities []Activity) (int, []Activity) {
	previousFinish := 0
	results := []Activity{}
	for _, currentActivity := range activities {
		if currentActivity.Start >= previousFinish {
			results = append(results, currentActivity)
			previousFinish = currentActivity.Finish
		}
	}
	return len(results), results
}
