package main

import "fmt"

func main() {

	c := make(chan int)

	// (비동기)
	go rangeSum(1000, c)
	go rangeSum(7000, c)
	go rangeSum(5000, c)

	// 순서대로 처리 (동기)
	result1 := <-c
	result2 := <-c
	result3 := <-c

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
}

func rangeSum(rg int, c chan int) {
	sum := 0

	for i := 0; i <= rg; i++ {
		sum += i
	}
	c <- sum
}
