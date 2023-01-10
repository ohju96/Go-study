package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Main : S --> ", time.Now())

	//var c chan int
	//c = make(chan int)

	v := make(chan int)
	go work1(v)
	go work2(v)

	<-v
	<-v
	fmt.Println("Main : S --> ", time.Now())
}

func work1(v chan int) {
	fmt.Println(" work 1 : S --> ", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println(" work 1 : E --> ", time.Now())
	v <- 1
}

func work2(v chan int) {
	fmt.Println(" work 2 : S --> ", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println(" work 2 : E --> ", time.Now())
	v <- 2
}
