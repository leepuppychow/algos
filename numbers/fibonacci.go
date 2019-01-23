package numbers

func FibonacciSlow(n int) int {
	if n <= 1 {
		return n
	}
	return FibonacciSlow(n-1) + FibonacciSlow(n-2)
}

func FibonacciFast(n int) int {
	fibList := []int{0, 1}
	for i := 2; i <= n; i++ {
		fibList = append(fibList, fibList[i-1]+fibList[i-2])
	}
	return fibList[n]
}