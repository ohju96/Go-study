package main

import "fmt"

var num int32

func init() {
	num = 30
	fmt.Println("Init Method start", num)
}

func main() {
	num = 40
	fmt.Println("Main Method start", num)
}
