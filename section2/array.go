package main

import "fmt"

func main() {
	var arr1 [5]int
	var arr2 [5]int = [5]int{1, 2, 3, 4, 5}
	var arr3 = [5]int{1, 2, 3, 4, 5}
	arr4 := [5]int{1, 2, 3, 4, 5}
	arr5 := [5]int{1, 2, 3}
	arr6 := [...]int{1, 2, 3, 4, 5}
	arr7 := [5][5]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9},
	}

	arr1[2] = 5

	fmt.Printf("%-5T %d %v\n", arr1, len(arr1), arr1)
	fmt.Printf("%-5T %d %v\n", arr2, len(arr2), arr2)
	fmt.Printf("%-5T %d %v\n", arr3, len(arr3), arr3)
	fmt.Printf("%-5T %d %v\n", arr4, len(arr4), arr4)
	fmt.Printf("%-5T %d %v\n", arr5, len(arr5), arr5)
	fmt.Printf("%-5T %d %v\n", arr6, len(arr6), arr6)
	fmt.Printf("%-5T %d %v\n", arr7, len(arr7), arr7)

	fmt.Printf("#####")

	arr8 := [5]int{
		1,
		2,
		3,
		4,
		5,
	}
	arr10 := [...]string{"oh", "ju", "hyeon"}
	fmt.Printf("%-5T %d %v\n", arr8, len(arr8), arr8)
	fmt.Printf("%-5T %d %v\n", arr10, len(arr10), arr10)

	fmt.Printf("#####")

	arr11 := [5]int{1, 10, 100, 1000, 10000}

	for i := 0; i < len(arr11); i++ {
		fmt.Println(arr11[i])
	}

	fmt.Printf("#####")

	arr22 := [5]int{7, 77, 777, 7777, 77777}

	for i, v := range arr22 {
		fmt.Println(i, v)
	}

	for _, v := range arr22 {
		fmt.Println(v)
	}

	fmt.Println("#####")

	for v := range arr22 {
		fmt.Println(v)
	}

	fmt.Println("#####")

	arr111 := [5]int{1, 10, 100, 1000, 10000}
	arr222 := arr111

	fmt.Println(arr111, &arr111)
	fmt.Println(arr222, &arr222)

	fmt.Printf("%p %v\n", &arr111, arr111)
	fmt.Printf("%p %v\n", &arr222, arr222)

}
