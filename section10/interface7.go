package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 타입 변환
	// 실행(런타임)
	var a interface{} = 1
	b := a
	c := a.(int)
	//d := a.(float64)

	fmt.Println(a)
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(b)
	fmt.Println(reflect.TypeOf(b))
	fmt.Println(c)
	fmt.Println(reflect.TypeOf(c))
	//fmt.Println(d)

	if v, ok := a.(int); ok {
		fmt.Println(v, ok)
	}
}
