package main

import "fmt"

func main() {
	map1 := map[string]int{
		"apple":  15,
		"banana": 115,
		"orange": 11115,
		"lemon":  0,
	}

	value1 := map1["lemon"]
	value2 := map1["kiwi"]
	value3, ok := map1["kiwi"]

	fmt.Println(value1)
	fmt.Println(value2)
	fmt.Println(value3, ok)

	fmt.Println()

	if value, ok := map1["kiwi"]; ok {
		fmt.Println(value)
	} else {
		fmt.Println("Kiwi is not exist !")
	}

	if value, ok := map1["banana"]; ok {
		fmt.Println(value)
	} else {
		fmt.Println("Kiwi is not exist !")
	}

	if _, ok := map1["kiwi"]; !ok {
		fmt.Println("Kiwi is not exist")
	}
}
