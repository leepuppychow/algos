// https://www.geeksforgeeks.org/greedy-algorithm-egyptian-fraction/

// Every positive fraction should be able to be broken down into sum of a unique set of unit fractions
// Ex: 2/3 = 1/2 + 1/3
// Ex: 6/14 = 1/3 + 1/11 + 1/231
// Ex: 12/13 is 1/2 + 1/3 + 1/12 + 1/156

package greedy

type Fraction struct {
	Numerator   int
	Denominator int
}

// Algorithm (greedy)
// 1. Start with current fraction and do D % N
// 		A. if D % N == 0 then add unit fraction (1 / (D/N))
// 		B. else if D % N != 0 then add unit fraction (1 / (1 + (D/N)))
// 2. Get current fraction - last unit fraction
// 3. Repeat step 1 for that remaining Fraction (break once D % N == 0)
