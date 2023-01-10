package main

import "fmt"

func main() {

	runFunc()

	fmt.Println("Hello Golang")
}

func runFunc() {

	defer func() {
		s := recover()
		fmt.Println("Error Message : ", s)
	}()

	panic("Error occurred!")
	fmt.Println("####")
}
