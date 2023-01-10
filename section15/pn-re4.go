package main

import (
	"fmt"
	"os"
)

func main() {

	fileopen("undefined.txt")
	fmt.Println("End main")
}

func fileopen(filename string) {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("File Open Error : ", r)
		}
	}()

	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("ex1 : ", f.Name())
	}

	defer f.Close()
}
