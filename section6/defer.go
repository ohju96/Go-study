package main

import "fmt"

func main() {
	ex_f1()
}

func ex_f1() {
	fmt.Println("f1 start")
	defer ex_f2()
	fmt.Println("f1 end")
}

func ex_f2() {
	fmt.Println("f2 start")
}
