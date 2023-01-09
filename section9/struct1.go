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

	oh := Account{number: "111-111", balance: 1000, interest: 0.2}
	lee := Account{number: "22-222", balance: 10}
	park := Account{number: "333-33", interest: 0.1}
	cho := Account{"444-44", 13000, 0.3}

	fmt.Println(oh)
	fmt.Println(lee)
	fmt.Println(park)
	fmt.Println(cho)

	fmt.Println()

	fmt.Println(int(oh.Calculate()))
	fmt.Println(int(lee.Calculate()))
	fmt.Println(int(park.Calculate()))
	fmt.Println(int(cho.Calculate()))

}
