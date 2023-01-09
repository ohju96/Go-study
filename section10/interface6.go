package main

import "fmt"

func main() {

	var a interface{}
	printContents(a)

	a = 7.5
	printContents(a)

	a = "Golang"
	printContents(a)

	a = true
	printContents(a)

	a = nil
	printContents(a)
}

func printContents(v interface{}) {
	fmt.Printf("Type : (%T) ", v)
	fmt.Println("ex1 : ", v)
}
