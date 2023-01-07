package main

import "fmt"

func main() {
	fmt.Println()

	var num1 int = 17
	var num2 int = -68
	var num3 int = 0631
	var num4 int = 0x32fa2c75

	fmt.Println(num1)
	fmt.Println(num2)
	fmt.Println(num3)
	fmt.Println(num4)

	fmt.Println("###")

	var char1 byte = 72
	var char2 byte = 0110
	var char3 byte = 0x48

	var char4 rune = 50556
	var char5 rune = 0142574
	var char6 rune = 0xc57c

	// %c %d %o %x

	fmt.Printf("%c %c %c\n", char1, char2, char3)
	fmt.Printf("%d %d %d\n", char1, char2, char3)
	fmt.Printf("%d %o %x\n", char1, char2, char3)

	fmt.Printf("%c %c %c\n", char4, char5, char6)
	fmt.Printf("%d %d %d\n", char4, char5, char6)
	fmt.Printf("%d %o %x\n", char4, char5, char6)

	fmt.Println("###")

	var num11 float32 = .14
	var num21 float32 = .75647
	var num31 float32 = 442.0378273
	var num41 float32 = 10.0

	var num5 float32 = 14e6
	var num6 float64 = .156875e+3
	var num7 float64 = 5.32521e-10

	fmt.Println(num11)
	fmt.Println(num21)
	fmt.Println(num31)
	fmt.Println(num41 - 0.1)
	fmt.Println(float32(num41 - 0.1))
	fmt.Println(float64(num41 - 0.1))
	fmt.Println(num5)
	fmt.Println(num6)
	fmt.Println(num7)
}
