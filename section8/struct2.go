package main

import "fmt"

func main() {

	a := cnt(10)
	var b cnt = 10

	fmt.Println(a)
	fmt.Println(b)

	fmt.Println()

	testConverT(int(b))
	testConverD(cnt(a))
}

type cnt int

func testConverT(i int) {
	fmt.Println(i)
}

func testConverD(i cnt) {
	fmt.Println(i)
}
