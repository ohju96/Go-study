package main

import "fmt"

func main() {
	// 타입 검사 switch 사용

	checkType(true)
	checkType(1)
	checkType(11.11)
	checkType("hi")
	checkType(nil)

}

func checkType(arg interface{}) {
	switch arg.(type) {
	case bool:
		fmt.Printf("This is a bool", arg)
	case int, int8, int16, int32, int64:
		fmt.Println("This is a int", arg)
	case float64:
		fmt.Println("This is a float", arg)
	case nil:
		fmt.Println("This is a nil", arg)
	default:
		fmt.Println("What is this type?")

	}

}
