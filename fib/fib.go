package fib

import "errors"

func fibRec(n int, memo map[int]int) int {
	if n < 3 {
		return 1
	}

	if _, key := memo[n]; key {
		return memo[n]
	}

	memo[n] = fibRec(n-1, memo) + fibRec(n-2, memo)

	return memo[n]
}

func FibRec(n int) (int, error) {
	if n < 1 {
		return 0, errors.New("N cannot be less than one")
	}

	return fibRec(n, make(map[int]int)), nil
}

func FibLoop(n int) (int, error) {
	if n < 1 {
		return 0, errors.New("N cannot be less than one")
	}

	a := []int{1, 1}
	var tmp int

	for i := 3; i <= n; i++ {
		tmp = a[1]
		a[1] += a[0]
		a[0] = tmp
	}

	return a[1], nil
}
