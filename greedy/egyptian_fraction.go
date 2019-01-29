// https://www.geeksforgeeks.org/greedy-algorithm-egyptian-fraction/

// Every positive fraction should be able to be broken down into sum of a unique set of unit fractions
// Ex: 2/3 = 1/2 + 1/3
// Ex: 6/14 = 1/3 + 1/11 + 1/231
// Ex: 12/13 is 1/2 + 1/3 + 1/12 + 1/156

package greedy

type Fraction struct {
	N int
	D int
}

var Fraction1 = Fraction{
	N: 2,
	D: 3,
}

var Fraction2 = Fraction{
	N: 6,
	D: 14,
}

var Fraction3 = Fraction{
	N: 30,
	D: 121,
}

// Recursive Algorithm (greedy)
// 1. Start with current fraction and do D % N
// 		A. if D % N == 0 then add unit fraction (1 / (D/N))
// 		B. else if D % N != 0 then add unit fraction (1 / (1 + (D/N)))
// 2. Get remaining difference = current fraction - last unit fraction
// 3. Repeat step 1 for that remaining Fraction (break once D % N == 0)

func EgyptianFraction(current Fraction, results []Fraction) (Fraction, []Fraction) {
	remainingFraction := Fraction{}
	unitFraction := Fraction{}
	if current.D%current.N == 0 {
		unitFraction = Fraction{
			N: 1,
			D: (current.D / current.N),
		}
		return unitFraction, append(results, unitFraction)
	} else {
		unitFraction = Fraction{
			N: 1,
			D: (current.D / current.N) + 1,
		}
		remainingFraction = Fraction{
			N: (current.N * unitFraction.D) - (current.D),
			D: current.D * unitFraction.D,
		}
		return EgyptianFraction(remainingFraction, append(results, unitFraction))
	}
}
