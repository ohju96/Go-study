package main

import "fmt"

func main() {

	for i := 0; i < 5; i++ {
		fmt.Println("ex1 : ", i)
	}

	// for {
	// 	fmt.Println("ex2 : Hello golang")
	// }

	loc := []string{"Seoule", "Busan", "Inchon"}

	for index, name := range loc {
		fmt.Println("ex3 : ", index, name)
	}
}
