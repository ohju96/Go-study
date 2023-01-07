package main

import (
	"fmt"
)

func main() {
	var b1 bool = true
	var b2 bool = false
	b3 := true

	fmt.Println(b1, b2, b3)
	fmt.Println(b1 == b2)
	fmt.Println(b1 && b2)
	fmt.Println(b1 || b2)
	fmt.Println(!b1)

	fmt.Println("###")

	fmt.Println(true && true)   //t
	fmt.Println(true && false)  //f
	fmt.Println(false && false) //f
	fmt.Println(false || false) //f
	fmt.Println(false || true)  //t
	fmt.Println(true || true)   //t
	fmt.Println(!false)         //t
	fmt.Println(!true)          //f

	fmt.Println("###")

	num1 := 15
	num2 := 37

	fmt.Println(num1 < num2)

}
