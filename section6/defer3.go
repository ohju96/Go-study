package main

import "fmt"

func main() {

	stack()
}

func stack() {
	fmt.Println("시작")

	for i := 1; i < 10; i++ {
		fmt.Println("for start")
		defer fmt.Println(i)
		fmt.Println("for end")
	}

	fmt.Println("종료")
}
