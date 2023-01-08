package main

import "fmt"

func main() {
	a()
}

func a() {
	defer end(start("b"))
	fmt.Println("in a")
}

func start(t string) string {
	fmt.Println("start :", t)
	return t
}

func end(t string) {
	fmt.Println("end : ", t)
}
