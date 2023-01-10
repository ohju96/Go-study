package main

import "fmt"

func main() {

	runFunc()

	fmt.Println("Hello Golang")
}

func runFunc() {

	defer func() {
		if s := recover(); s != nil {
			fmt.Println("Error Message : ", s)
		}
	}()

	a := [3]int{1, 2, 3}

	for i := 0; i < 5; i++ {
		fmt.Println("ex 1 : ", a[i])

	}
}
