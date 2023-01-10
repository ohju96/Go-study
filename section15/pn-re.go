package main

import "fmt"

func main() {
	// 패닉 -> 디퍼 -> 리커버 순서로 진행

	fmt.Println("Start Main")
	panic("Error user stopped!")
	fmt.Println("End Main")

}
