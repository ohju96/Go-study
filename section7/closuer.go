package main

import "fmt"

func main() {

	multiplay := func(x, y int) int {
		return x * y
	}

	r1 := multiplay(5, 10)

	fmt.Println(r1)

	fmt.Println()

	m, n := 5, 10
	sum := func(c int) int {
		return n + m + c
	}

	r2 := sum(10)
	fmt.Println(r2)
}
