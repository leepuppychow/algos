package arrays

import (
	"fmt"
)

type Solution struct {
	Num1   int
	Index1 int
	Num2   int
	Index2 int
}

func SumOfPairInArray(sum int, arr []int) []Solution {
	fmt.Println("Find if there is a pair with given sum in an array")
	hash := make(map[int]int)
	var solutions []Solution
	var s Solution

	for i, number := range arr {
		hash[number] = i

		i2, ok := hash[sum-number]
		if ok {
			s = Solution{
				Num1:   number,
				Index1: i,
				Num2:   sum - number,
				Index2: i2,
			}
			solutions = append(solutions, s)
		}
	}
	return solutions
}
