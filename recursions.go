package main
func factorial(n int) (fac int) {
	if n <=1 {
		fac = 1
	} else {
		fac = n*factorial(n-1)
	}
	return
}
func fibonacci(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n - 1) + fibonacci(n - 2)
	}
	return
}