package main

import "fmt"

type Employee struct {
	name   string
	salary float64
	bonus  float64
}

// 구조체 임베디드 패턴
type Executives struct { // 구조체
	Employee     // is a 관계 : 임원도 직원이다.
	specialBonus float64
}

func (e Employee) Calculate() float64 {
	return e.salary + e.bonus
}

func main() {

	ep1 := Employee{"oh", 100, 10}
	ep2 := Employee{"kim", 200, 20}
	ex := Executives{Employee{"lee", 300, 30}, 30}

	fmt.Println(ep1.Calculate())
	fmt.Println(ep2.Calculate())

	fmt.Println(ex.Calculate() + ex.specialBonus)
}
