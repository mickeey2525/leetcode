package main

import "fmt"

var cache = map[int]int{0: 0, 1: 1}

func fib(N int) int {
	if N <= 1 {
		return N
	}
	return memoize(N)
}

func memoize(N int) int {
	if _, ok := cache[N]; ok {
		return cache[N]
	}
	cache[N] = memoize(N-1) + memoize(N-2)
	return memoize(N)
}

func main() {
	a :=fib(100)
	fmt.Printf("The result is %d", a)
}
