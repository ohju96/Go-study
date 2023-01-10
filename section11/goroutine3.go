package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println("Current System CPU Start :", runtime.GOMAXPROCS(0)) // 설정 값 뽑기

	fmt.Println("Main Routine Start : ", time.Now())

	for i := 0; i < 10; i++ {
		go exe(i)
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Current System CPU End :", runtime.GOMAXPROCS(0))

}

func exe(name int) {
	r := rand.Intn(100)
	fmt.Println(name, " start : ", time.Now())
	for i := 0; i < 100; i++ {
		fmt.Println(name, " >>>>> ", r, i)
	}
	fmt.Println(name, " end : ", time.Now())
}
