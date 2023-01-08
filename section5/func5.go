package main

import "fmt"

func main() {
	f := []func(int, int) int{multiply5, sum5}
	a := f[0](10, 10)
	b := f[1](10, 10)

	fmt.Println(a)
	fmt.Println(b)

	var f1 func(int, int) int = multiply5
	f2 := sum5

	fmt.Println(f1(10, 10))
	fmt.Println(f2(10, 10))
	fmt.Println()

	m := map[string]func(int, int) int{
		"mulFunc": multiply5,
		"sumFunc": sum5,
	}

	fmt.Println(m["mulFunc"](10, 10))
	fmt.Println(m["sumFunc"](10, 10))

}

func multiply5(x, y int) (r int) {
	r = x * y
	return r
}

func sum5(x, y int) (r int) {
	r = x + y
	return r
}
