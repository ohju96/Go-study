package main

import "fmt"

func main() {

	a, b := multiply(10, 5)
	c, _ := multiply(10, 5)
	_, d := multiply(10, 5)

	fmt.Println(a, b)
	fmt.Println(c)
	fmt.Println(d)

	fmt.Println()

	x1, x2, x3, x4, x5 := arrayMultiply(1, 2, 3, 4, 5)
	y1, _, _, _, y5 := arrayMultiply(1, 2, 3, 4, 5)

	fmt.Println(x1, x2, x3, x4, x5)
	fmt.Println(y1, y5)

	fmt.Println()

	r1, r2 := multiplyResult(1, 2)
	fmt.Println(r1, r2)
}

func multiply(x int, y int) (int, int) {
	return x * 10, y * 10
}

func arrayMultiply(a, b, c, d, e int) (int, int, int, int, int) {
	return a * 1, b * 2, c * 3, d * 4, e * 5
}

func multiplyResult(x int, y int) (r1 int, r2 int) {
	r1 = x * 10
	r2 = y + 10
	return r1, r2
}
