package main

import "fmt"

var logf = fmt.Printf
var log = fmt.Println

// recursive
func recurFac(n int) int {
	logf("%d * ", n)
	if n == 1 {
		return 1
	}
	return n * recurFac(n-1)
}

// tail recursion
func tailRecurFac(count int, accum int) int {
	if count == 0 {
		return accum
	} else {
		return tailRecurFac(count-1, count*accum)
	}
}
func tailFac(n int) int {
	return tailRecurFac(n, 1)
}

// tail recursion
func loopFac(count int, accum int) int {
	for {
		if count == 0 {
			return accum
		} else {
			accum = count * accum
			count -= 1
			continue
			// return tailRecurFac(count-1, count*accum)
		}
	}
}
func LoopFac(n int) int {
	return loopFac(n, 1)
}

func main() {

	// recurFac(10) // recursive
	// log(tailFac(4))
	log(LoopFac(4))
}
