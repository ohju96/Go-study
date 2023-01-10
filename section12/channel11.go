package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan string)

	go func() {
		for {
			num := <-ch1
			fmt.Println("ch1 : ", num)
			time.Sleep(250 * time.Millisecond)
		}
	}()

	go func() {
		ch2 <- "Golang Hi!"
		time.Sleep(500 * time.Millisecond)
	}()

	go func() {
		for {
			select {
			case ch1 <- 777:
			case str := <-ch2:
				fmt.Println("ch2 : ", str)
			}
		}
	}()

	time.Sleep(1 * time.Second)
}
