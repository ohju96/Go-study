package main

import "fmt"

func main() {

	i := 7
	p := &i
	fmt.Println(i, *p, &i, p)

	*p++

	fmt.Println(i, *p, &i, p)

	*p = 7777

	fmt.Println(i, *p, &i, p)

	i = 77

	fmt.Println(i, *p, &i, p)

}
