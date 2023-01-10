package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	a, err := notZero(1)

	if err != nil {
		// log.Fatal(err)
		log.Fatal(err.Error())
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
	return "", errors.New("0을 입력했습니다. 에러 발생 !")
}
