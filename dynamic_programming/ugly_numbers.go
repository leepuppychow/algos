// https://www.geeksforgeeks.org/ugly-numbers/

// Ugly (or Hamming) number is one that only has 2,3, or 5 as its prime factors
// Given n, find the nth Ugly number in the sequence of ugly numbers:
// Ex:  1, 2, 3, 4, 5, 6, 8, 9, 10, 12, 15, …

// RANDOM NOTES ON DYNAMIC PROGRAMMING

// A problem suited for DP approach satisfy:
// 1. Overlapping subproblems: as you go through an algorithm certain values are being calculated repeatedly.
// 		Instead of wasting that effort of re-calculating, store that value somehow (tabulation of memoization)
// 2. Optimal substructure property: meaning that the global optimum can be solved by using the local optimum of each subproblem.

// --> Generally need to decide what state to store. This state is like an ID for each subproblem.
// --> Then try to find the relationship between these states.
// --> Store this state via tabulation (bottom-up) vs memoization (top-down)

package dp

import (
	"fmt"
)

// Better algorithm (dynamic programming approach)
// Key idea: why not split the sequence into 3 subsets (multiples of 2,3,5 * values of ugly num series)

// 1  2  3  4  5   6   8    9  10  12...
// 2  4  6  8  10  12  16  18  20  24...
// 3  6  9  12 15  18  24  27  30  36...
// 5  10 15 20 25  50  40  45  50  60...

// Find the current minimum out of those subsets and then add that to the series
// Now, you do not have to do brute force approach iterating by 1 and checking if it is an Ugly number

// ALGORITHM:
/*

	uglyNums: [1, 2, 3, 4, 5, 6, 8, 9, 10 ... ]
	i2  i3    i5  mult2  mult3  mult5
	0    0    0    *2      3      5
	1    0    0     4     *3      5
	1    1    0    *4      6      5
	2    1    0     6      6     *5
	2    1    1    *6     *6      10
	3    2    1    *8      9      10
	4    2    1     10    *9      10
	4    3    1    *10     12    *10
	5    3    2    *12    *12     15
							.
							.
							.
*/
// This should be O(n) for time and space

func UglyNumber(n int) int {
	uglyNums := []int{1}
	i2, i3, i5 := 0, 0, 0

	for i := 1; i <= n; i++ {
		mult2 := uglyNums[i2] * 2
		mult3 := uglyNums[i3] * 3
		mult5 := uglyNums[i5] * 5
		min := Min(mult2, mult3, mult5)

		if min == mult2 {
			i2++
		}
		if min == mult3 {
			i3++
		}
		if min == mult5 {
			i5++
		}

		uglyNums = append(uglyNums, min)
	}
	return uglyNums[n-1]
}

func Min(nums ...int) int {
	min := nums[0]
	for _, num := range nums {
		if num < min {
			min = num
		}
	}
	return min
}

// If a number is divisible by a combination of existing Hamming numbers in our collection, then it is a Hamming number.
// Maybe create Hash table of Hamming number => True

// 1. start with Hash (first 5 entries)
// 2. for i = 6 up to n
// 3. if (hash[n/2] || hash[n/3] || hash[n/5]) then it must be a Hamming number
// 4. Add it to the Hash, and iterate

func SlowUglyNumber(n int) int {
	var mostRecentNum int
	hamHash := map[int]bool{
		1: true,
		2: true,
		3: true,
		4: true,
		5: true,
	}
	num := 6

	for len(hamHash) < n {
		ok2, ok3, ok5 := false, false, false
		if num%2 == 0 {
			if _, ok2 = hamHash[num/2]; ok2 {
				hamHash[num] = true
				mostRecentNum = num
			}
		} else if num%3 == 0 {
			if _, ok3 = hamHash[num/3]; ok3 {
				hamHash[num] = true
				mostRecentNum = num
			}
		} else if num%5 == 0 {
			if _, ok5 = hamHash[num/5]; ok5 {
				hamHash[num] = true
				mostRecentNum = num
			}
		}
		num++ // THIS WILL BLOW UP EXPONENTIALLY...
	}
	fmt.Println(mostRecentNum)
	return mostRecentNum
}
