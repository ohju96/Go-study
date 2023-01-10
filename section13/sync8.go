package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	var cnt int64 = 0
	wg := new(sync.WaitGroup)

	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func(n int) {
			cnt += 1
			fmt.Println("WaitGroup : ", n)
			wg.Done()
		}(i)
	}

	for i := 0; i < 2000; i++ {
		wg.Add(1)
		go func(n int) {
			cnt -= 1
			fmt.Println("WaitGroup : ", n)
			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Println("WaitGroup End ! cnt? >>>>> ", cnt)
}
