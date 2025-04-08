package fibonacci

// Fibonacci that returns a function that returns an int
func Fibonacci(n int) int {
	if n <= 0 {
		return 0
	} else if n == 1 {
		return 0
	} else if n == 2 {
		return 1
	}

	a, b := 0, 1
	for i := 3; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}
