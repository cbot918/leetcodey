package main

import "fmt"

func main() {
	q := []int{}

	q = append(q, 1)
	q = append(q, 3)

	x := q[0]
	fmt.Println(x)
	q = q[1:]
	x = q[0]
	fmt.Println(x)
}
