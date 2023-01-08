package main

import "fmt"

func main() {

	x := multiply4(1, 2, 3, 4, 5)
	fmt.Println(x)
	fmt.Println()

	y := sum4(10, 2, 3, 1, 4)
	fmt.Println(y)
	fmt.Println()

	prtWorld("a", "b", "test")

	fmt.Println()
	a := []int{1, 2, 3, 4, 5, 6}
	m := multiply4(a...)
	i := sum4(a...)

	fmt.Println(m)
	fmt.Println(i)

}

func prtWorld(msg ...string) {
	for _, value := range msg {
		fmt.Println(value)
	}
}

func sum4(n ...int) int {
	tot := 1
	for _, value := range n {
		tot += value
	}
	return tot
}

func multiply4(n ...int) int {
	tot := 1
	for _, value := range n {
		tot += value
	}
	return tot
}
