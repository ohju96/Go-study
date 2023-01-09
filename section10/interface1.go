package main

import "fmt"

func main() {

	var t test
	fmt.Println(t) //빈 경우 nil
}

type test interface{}
