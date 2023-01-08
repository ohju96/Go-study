package main

import "fmt"

func main() {
	x := fact(7)
	fmt.Println(x)

	prtHello(10)

}

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func prtHello(n int) {
	if n == 0 {
		return
	}
	fmt.Println("(", n, ")", "hi")
	prtHello(n - 1)
}
