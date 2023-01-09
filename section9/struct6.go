package main

import (
	"fmt"
)

type Car1 struct {
	name    string "차량명"
	color   string "색상"
	company string "제조사"
	detail  spec   "상세"
}

type spec struct {
	length int "전장"
	height int "전고"
	width  int "전축"
}

func main() {

	car1 := Car1{
		"520d",
		"silver",
		"bmw",
		spec{39, 2, 3},
	}

	fmt.Println(car1.name)
	fmt.Println(car1.color)
	fmt.Println(car1.company)
	fmt.Printf("%#v\n", car1.detail)

	fmt.Println()

	fmt.Println(car1.detail.length)
	fmt.Println(car1.detail.height)
	fmt.Println(car1.detail.width)

}
