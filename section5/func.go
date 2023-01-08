package main

import (
	"fmt"
	"strconv"
)

func main() {
	// func함수명(매개변수)(반환타입):반환값 리턴 값 여러 개 가능하다.

	helloGolang()
	say_one("Hi, there")

	result := sum1(5, 5)

	resutlString := strconv.Itoa(result)
	fmt.Println(result)
	fmt.Println(resutlString) // int -> string
}

func helloGolang() {
	fmt.Println("Hello Golang")
}

func say_one(m string) {
	fmt.Println(m)
}
func sum1(x int, y int) int {
	return x + y
}
