package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	f, err := os.Open("unnamedfile") //err
	if err != nil {
		// log.Fatal(err.Error())
		log.Fatal(err)
	}

	fmt.Println("======")
	fmt.Println(f.Name())
}
