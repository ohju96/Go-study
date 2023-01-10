package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	var cnt int64 = 0
	wg := new(sync.WaitGroup)

	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func(n int) {
			//cnt += 1
			atomic.AddInt64(&cnt, 1)
			fmt.Println("WaitGroup : ", n)
			wg.Done()
		}(i)
	}

	for i := 0; i < 2000; i++ {
		wg.Add(1)
		go func(n int) {
			//cnt -= 1
			atomic.AddInt64(&cnt, -1)
			fmt.Println("WaitGroup : ", n)
			wg.Done()
		}(i)
	}

	wg.Wait()
	finalCnt := atomic.LoadInt64(&cnt)
	fmt.Println("WaitGroup End ! cnt? >>>>> ", cnt)
	fmt.Println("WaitGroup End ! cnt? >>>>> ", finalCnt)
}
