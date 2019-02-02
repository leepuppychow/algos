package dp

import "fmt"

// https://www.geeksforgeeks.org/bell-numbers-number-of-ways-to-partition-a-set/

// Bell number is the number of ways a set of (n) elements can be partitioned.
/*
	Examples:
		n = 1 --> 1
		n = 2 --> 2
		n = 3 --> 5
		n = 4 --> 15
		n = 5 --> 52
		n = 6 --> 203
*/

// Algorithm: Bell triangle (how to derive this??)

// 1
// 1  2
// 2  3  5
// 5  7  10  15
// 15 20 27  37  52
// 52 67 87 114 151  203

// 1. Set currentRow[0] = last item in previousRow
// 2. Append to currentRow with currentRow[i-1]+previousRow[i-1]
// 3. Break out of the recursion once the previousRow length == n (means we are on the currentRow we want)

func BellNumber(n int, previousRow []int) (int, []int) {
	if len(previousRow) == n {
		return previousRow[n-1], previousRow
	}
	var currentRow []int
	currentRow = append(currentRow, previousRow[len(previousRow)-1])
	for i := 1; i <= len(previousRow); i++ {
		currentRow = append(currentRow, currentRow[i-1]+previousRow[i-1])
	}
	fmt.Println(currentRow)
	return BellNumber(n, currentRow)
}
