package recursion

func simplest_case(n int) int {
	if n == 1 {
		return 1
	}

	return n + simplest_case(n-1)
}

func fib(n int) int {

	if n == 0 || n == 1 {
		return n
	}

	return fib(n-1) + fib(n-2)
}
