package main

import "fmt"

func main() {
	var a *int
	var b *int = new(int)
	fmt.Println()

	fmt.Println(a)
	fmt.Println(b)

	fmt.Println()

	i := 7
	a = &i
	b = &i

	fmt.Println(a)
	fmt.Println(&a)
	fmt.Println(*a) // 역참조

	fmt.Println(b)
	fmt.Println(&b)
	fmt.Println(*b)

	var c = &i
	d := &i

	*d = 10

	fmt.Println()

	fmt.Println(*c)
	fmt.Println(*d)
}
