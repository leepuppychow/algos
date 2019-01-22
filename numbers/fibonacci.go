package numbers

func FibonacciSlow(n int) int {
	if n <= 1 {
		return n
	}
	return FibonacciSlow(n-1) + FibonacciSlow(n-2)
}
