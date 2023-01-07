package main

import (
	"fmt"
	"sort"
)

func main() {
	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(slice1[:])
	fmt.Println(slice1[0:])
	fmt.Println(slice1[:5])
	fmt.Println(slice1[0:len(slice1)])
	fmt.Println(slice1[3:])
	fmt.Println(slice1[:3])
	fmt.Println(slice1[1:3]) // [이상:미만]

	fmt.Println()

	// https://golang.org/pkg/sort

	slice2 := []int{3, 6, 10, 9, 4, 5, 8, 3, 4}
	slice3 := []string{"가", "다", "b", "z", "i"}

	fmt.Println(sort.IntsAreSorted(slice2)) //정렬확인
	sort.Ints(slice2)                       //정렬
	fmt.Println(slice2)

	fmt.Println()

	fmt.Println(sort.StringsAreSorted(slice3))
	sort.Strings(slice3)
	fmt.Println(sort.StringsAreSorted(slice3))
}
