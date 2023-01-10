package main

import (
	"fmt"
	"time"
)

func main() {

	exe("t1") // 먼저 다 찍힘
	fmt.Println("Main Routine Start : ", time.Now())
	go exe("t2")
	go exe("t3")
	time.Sleep(4 * time.Second)
	fmt.Println("Main Routine end : ", time.Now())
}

func exe(name string) {
	fmt.Println(name, " start : ", time.Now())
	for i := 0; i < 100; i++ {
		fmt.Println(name, ">>>>>>", i)
	}
	fmt.Println(name, " end : ", time.Now())
}
