package main

import "fmt"

func main() {
	cnt := increaseCnt()

	fmt.Println(cnt())
	fmt.Println(cnt())
	fmt.Println(cnt())
	fmt.Println(cnt())
	fmt.Println(cnt())
	fmt.Println(cnt())
	fmt.Println(cnt())
	fmt.Println(cnt())
	fmt.Println(cnt())
	fmt.Println(cnt())
}

func increaseCnt() func() int {
	var n int = 0
	fmt.Println(n)
	return func() int {
		fmt.Println("START")
		n += 1
		return n
	}
}
