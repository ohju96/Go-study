package main

import "fmt"

func main() {

	func() {
		fmt.Println("func")
	}()

	msg := "hello"

	func(m string) {
		fmt.Println(m)
	}(msg)

	func(x, y int) {
		fmt.Println(x + y)
	}(10, 20)

	r := func(x, y int) int {
		return x * y
	}(10, 100)

	fmt.Println(r)
}
