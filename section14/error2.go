package main

import (
	"fmt"
	"log"
)

func main() {
	a, err := notZero(1)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(a)

	b, err := notZero(0)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(b)
}

func notZero(n int) (string, error) {
	if n != 0 {
		s := fmt.Sprint("Hello", n)
		return s, nil
	}
	return "", fmt.Errorf("%d를 입력했습니다. 에러 발생 !", n)
}
