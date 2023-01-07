package main

import (
	"fmt"
)

func main() {

	a := 30 / 15
	switch a {
	case 2, 4, 6:
		fmt.Println("a -> ", a, "는 짝수")
	case 1, 3, 5:
		fmt.Println("a -> ", a, "는 홀수")
	}

	switch e := "go"; e {
	case "java":
		fmt.Println("자바다")
	case "go":
		fmt.Println("고다")
		fallthrough
	case "python":
		fmt.Println("파이썬이다")
	case "ruby":
		fmt.Println("루비다")
		fallthrough
	case "c":
		fmt.Println("씨다")

	}
}
