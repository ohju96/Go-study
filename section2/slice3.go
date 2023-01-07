package main

import "fmt"

func main() {

	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{8, 9, 10, 11, 12}
	s3 := []int{13, 14, 15, 16, 17}

	s1 = append(s1, 6, 7)
	s2 = append(s1, s2...)
	s3 = append(s2, s3[0:3]...)

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)

	fmt.Println()

	s4 := make([]int, 0, 5)
	for i := 0; i < 15; i++ {
		s4 = append(s4, i)
		fmt.Printf("p : %p, len : %d, cap : %d, value : %v\n", &s4, len(s4), cap(s4), s4)
	}

}
