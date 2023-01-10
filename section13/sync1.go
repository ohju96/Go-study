package main

import (
	"fmt"
	"runtime"
)

type count struct {
	num int
}

func (c *count) increment() {
	c.num += 1
}

func (c *count) result() {
	fmt.Println(c.num)
}

func main() {
	// 고루틴 동기화 예제
	// 실행 흐름 제어 및 변수 동기화 가능
	// 공유 데이터 보호가 가장 중요 (스레드 세이프.. -> 동기화 관련)

	runtime.GOMAXPROCS(runtime.NumCPU()) //시스템 전체 CPU 사용

	c := count{num: 0}
	done := make(chan bool)

	for i := 1; i <= 10000; i++ {
		go func() {
			c.increment()
			done <- true
			runtime.Gosched() //양보
		}()
	}

	for i := 1; i <= 10000; i++ {
		<-done
	}

	c.result()

}
