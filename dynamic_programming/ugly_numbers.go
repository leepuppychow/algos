// https://www.geeksforgeeks.org/ugly-numbers/

// Ugly (or Hamming) number is one that only has 2,3, or 5 as its prime factors
// Given n, find the nth Ugly number in the sequence of ugly numbers:
// Ex:  1, 2, 3, 4, 5, 6, 8, 9, 10, 12, 15, …

// If a number is divisible by a combination of existing Hamming numbers in our collection, then it is a Hamming number.
// Maybe create Hash table of Hamming number => True

// 1. start with Hash (first 5 entries)
// 2. for i = 6 up to n
// 3. if (hash[n/2] || hash[n/3] || hash[n/5]) then it must be a Hamming number
// 4. Add it to the Hash, and iterate

package dp

import "fmt"

func UglyNumber(n int) int {
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
