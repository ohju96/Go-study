package main

import (
	"fmt"
	"math"
)

func main() {

	var n1 uint8 = math.MaxUint8
	var n2 uint16 = math.MaxUint16
	var n3 uint32 = math.MaxUint32
	var n4 uint64 = math.MaxUint64

	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println(n3)
	fmt.Println(n4)

	fmt.Println("###")

	n5 := 100000
	n6 := int16(10000)
	n7 := uint8(100)

	fmt.Println(n5 + int(n6))
	fmt.Println(int16(n7) + n6)
}
