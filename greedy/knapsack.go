// https://www.geeksforgeeks.org/fractional-knapsack-problem/

package greedy

type Item struct {
	Id               int
	Value            int
	Weight           int
	ValueWeightRatio int
}

var Knapsack = []Item{
	Item{
		Id:               1,
		Value:            120,
		Weight:           30,
		ValueWeightRatio: 0,
	},
	Item{
		Id:               2,
		Value:            100,
		Weight:           20,
		ValueWeightRatio: 0,
	},
	Item{
		Id:               3,
		Value:            60,
		Weight:           10,
		ValueWeightRatio: 0,
	},
}

func GetValueWeightRatios(knapsack []Item) []Item {
	result := []Item{}
	for _, item := range knapsack {
		item.ValueWeightRatio = item.Value / item.Weight
		result = append(result, item)
	}
	return result
}

func SortByRatio(knapsack []Item) []Item {
	// Bubble Sort (could pick a more efficient for refactoring...)
	// O(n^2)
	for i := 0; i < len(knapsack); i++ {
		for i := 0; i < len(knapsack)-1; i++ {
			if knapsack[i].ValueWeightRatio < knapsack[i+1].ValueWeightRatio {
				knapsack[i], knapsack[i+1] = knapsack[i+1], knapsack[i]
			}
		}
	}
	return knapsack
}

func FractionalKnapsack(knapsack []Item, weightLimit int) int {
	maxValue := 0
	knapsack = SortByRatio(GetValueWeightRatios(knapsack)) // Rate limiting step O(n^2)
	for _, item := range knapsack {
		if weightLimit == 0 {
			return maxValue
		} else if item.Weight <= weightLimit { // Since they are sorted by ratio, pick the most valuable items first
			maxValue += item.Value
			weightLimit -= item.Weight
		} else { // Now add in the remaining value for the most valuable item at this point in time
			maxValue += weightLimit * item.ValueWeightRatio
		}
	}
	return maxValue
}
