package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := new(sync.WaitGroup)

	for i := 0; i < 2; i++ {

		wg.Add(1)

		go func(n int) {
			fmt.Println("WaitGroup : ", n)
			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Println("WaitGroup End !")
}
