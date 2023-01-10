package main

import (
	"errors"
	"fmt"
)

type err interface {
	Error() string
}

func main() {
	var err1 error = errors.New("Error occurred -1")
	err2 := errors.New("Error occurred -2")

	fmt.Println(err1)
	fmt.Println(err1.Error())

	fmt.Println(err2)
	fmt.Println(err2.Error())

}
