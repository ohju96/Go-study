package main

import "fmt"

func rptc(n *int) {
	*n = 77
}

func vptc(n int) {
	n = 77
}

func main() {

	var a int = 10
	var b int = 10

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println()

	rptc(&a)
	vptc(b)

	fmt.Println(a)
	fmt.Println(b)
}
