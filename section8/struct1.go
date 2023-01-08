package main

import "fmt"

func main() {
	bmw := Car{"520d", "Red", 10, 1}
	benz := Car{"220d", "blue", 1, 10}

	fmt.Println(bmw, &bmw)
	fmt.Println(benz, &benz)

	fmt.Println()

	fmt.Println(Price(bmw))
	fmt.Println(Price(benz))

	fmt.Println()

	fmt.Println(bmw.Price())
	fmt.Println(benz.Price())
}

type Car struct {
	name   string
	coloer string
	price  int64
	tax    int64
}

// 일반 메서드
func Price(c Car) int64 {
	return c.price + c.tax
}

// 일반 + 구조체 바인딩 메서드
func (c Car) Price() int64 {
	return c.price + c.tax

}
