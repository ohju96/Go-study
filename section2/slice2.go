package main

import "fmt"

func main() {

	arr1 := [3]int{1, 2, 3}
	var arr2 [3]int

	arr2 = arr1
	arr2[0] = 7

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println()

	slice1 := []int{1, 2, 3}
	var slice2 []int
	slice2 = slice1
	slice2[0] = 7

	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println()

	slice3 := make([]int, 50, 100)
	fmt.Println(slice3[40])

	for i, v := range slice1 {
		fmt.Println(i, v)
	}

	fmt.Println()

	test := []int{0, 1, 2, 3, 4}       // test 슬라이스에 int형식으로 데이터를 넣는다.
	copyTest := make([]int, len(test)) // opyTest 슬라이스에 길이를 test와 같은 슬라이스를 만들어준다.

	copy(copyTest, test) // test를 copyTest에 복사해 넣어준다.

	copyTest[0] = 100 // copyTest 인덱스 0에 100을 넣어준다.

	fmt.Printf("%p : %s \n", &copyTest, copyTest)
	fmt.Printf("%p : %s \n", &test, test)

}
