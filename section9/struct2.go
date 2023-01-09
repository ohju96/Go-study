package main

import "fmt"

type Account struct {
	number   string
	balance  float64
	interest float64
}

func (a Account) Calculate() float64 {
	return a.balance + (a.balance * a.interest)
}

func main() {

	var kim *Account = new(Account)
	kim.number = "12-12"
	kim.balance = 100
	kim.interest = 0.1

	oh := &Account{"12", 10, 1}
	lee := new(Account)
	lee.number = "1"
	lee.balance = 11
	lee.interest = 9.2

	fmt.Println(kim)
	fmt.Println(oh)
	fmt.Println(lee)

	fmt.Printf("ex : %#v\n", kim)
	fmt.Printf("ex : %#v\n", oh)
	fmt.Printf("ex : %#v\n", lee)

	fmt.Println()

	fmt.Println("ex3 : ", int(kim.Calculate()))
	fmt.Println("ex3 : ", int(oh.Calculate()))
	fmt.Println("ex3 : ", int(lee.Calculate()))
}
