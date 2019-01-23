// Greatest common divisor (using Euclidean algorithm)

package numbers

import "fmt"

func GCD(a, b int) int {
	// Just assume that a is the larger number
	fmt.Println("GCD called")
	if b == 0 {
		return a
	}
	return GCD(b, a%b)
}
