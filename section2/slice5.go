package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3, 4, 5, 6, 7}
	slice2 := make([]int, 5)
	slice3 := []int{}

	copy(slice2, slice1)
	copy(slice3, slice1)

	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3) //copy를 사용하기 위해 make로 공간 할당을 해야한다.

	fmt.Println()

	a := []int{1, 2, 3, 4, 5}
	b := make([]int, 5)

	copy(b, a)

	b[0] = 7
	b[4] = 10
	a[3] = 11

	fmt.Println(a)
	fmt.Println(b)

	fmt.Println()

	c := [5]int{1, 2, 3, 4, 5}
	d := c[0:3]

	d[1] = 7
	c[1] = 11

	fmt.Println(c)
	fmt.Println(d)

	fmt.Println()

	e := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	f := e[0:5:7]

	fmt.Println(len(f), cap(f))
	fmt.Println(f)
}
