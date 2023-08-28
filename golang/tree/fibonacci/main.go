package main

import "fmt"

var logf = fmt.Printf

// recursive with cache

var cache map[int]int

func fcache(in int) int {
	if val, found := cache[in]; found {
		return val
	}
	if in < 2 {
		cache[in] = in
	} else {
		cache[in] = fcache(in-1) + fcache(in-2)
	}
	return cache[in]
}

// recursive
func f(in int) int {
	if in == 0 {
		return 0
	}
	if in == 1 {
		return 1
	}
	return f(in-1) + f(in-2)
}

// tail recursive
func fibonacciRecursively(counter int, init int, accumulator int) int {
	if counter == 0 {
		return init
	} else {
		return fibonacciRecursively(counter-1, accumulator, init+accumulator)
	}
}
func Fibonacci(n int) int {
	return fibonacciRecursively(n, 0, 1)
}

// loop version
func loopFibo(counter int, beadd int, add int) int {
	for {
		if counter == 0 {
			return beadd
		} else {
			counter -= 1
			temp := add
			add = add + beadd
			beadd = temp
			continue
		}
	}
}

func Fibo(n int) int {
	return loopFibo(n, 0, 1)
}

func main() {
	max := 10
	// flag := "recursive"
	// flag := "tail recursive"
	flag := "fcache"
	cache = make(map[int]int)
	for i := 0; i < max; i++ {
		if flag == "recursive" {
			logf("%d ", f(i))
		}
		if flag == "tail recursive" {
			logf("%d ", Fibonacci(i))
		}
		if flag == "loop version" {
			logf("%d ", Fibo(i))
		}
		if flag == "fcache" {
			logf("%d ", fcache(i))
		}
	}
}
