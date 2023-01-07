package main

import "fmt"

func main() {

	var slice1 []int
	slice2 := []int{}
	slice3 := []int{1, 2, 3, 4, 5}
	slice4 := [][]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9},
	}

	slice3[4] = 10

	fmt.Printf("%-5T %d %d %v\n", slice1, len(slice1), cap(slice1), slice1)
	fmt.Printf("%-5T %d %d %v\n", slice2, len(slice2), cap(slice2), slice2)
	fmt.Printf("%-5T %d %d %v\n", slice3, len(slice3), cap(slice3), slice3)
	fmt.Printf("%-5T %d %d %v\n", slice4, len(slice4), cap(slice4), slice4)

	fmt.Println()

	// 길이와 용량을 정해서 사용할 수 있다.
	var slice5 []int = make([]int, 10, 10)
	var slice6 = make([]int, 5)
	slice7 := make([]int, 5, 100)
	slice8 := make([]int, 8)

	slice6[2] = 7

	fmt.Printf("%-5T %d %d %v\n", slice1, len(slice5), cap(slice5), slice5)
	fmt.Printf("%-5T %d %d %v\n", slice2, len(slice6), cap(slice6), slice6)
	fmt.Printf("%-5T %d %d %v\n", slice3, len(slice7), cap(slice7), slice7)
	fmt.Printf("%-5T %d %d %v\n", slice4, len(slice8), cap(slice8), slice8)

	fmt.Println()

	var slice9 []int

	if slice9 == nil {
		fmt.Println("This is Nil Slice!")
	}

}
