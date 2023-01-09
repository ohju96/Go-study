package main

import (
	"fmt"
	"time"
)

func main() {
	exe1()

	fmt.Println("Main Routine Start", time.Now())
	go exe2()
	go exe3()
	time.Sleep(4 * time.Second)
	fmt.Println("Main Routine end", time.Now())

}

func exe1() {
	fmt.Println("exe1 func start", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("exe1 func end", time.Now())
}

func exe2() {
	fmt.Println("exe2 func start", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("exe2 func end", time.Now())
}

func exe3() {
	fmt.Println("exe3 func start", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("exe3 func end", time.Now())
}
