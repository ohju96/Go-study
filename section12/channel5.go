package main

import (
	"fmt"
)

func main() {

	ch := make(chan bool)

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("check : ", i)
			ch <- true
		}
		close(ch) // 5회 채널에 값 전송 후 채널 닫기
	}()

	for i := range ch { // 채널에서 값을 꺼내온다. 닫힐 때 까지
		fmt.Println("ex1 : ", i)
	}
}
