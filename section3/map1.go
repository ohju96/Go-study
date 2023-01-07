package main

import "fmt"

func main() {

	var map1 map[string]int = make(map[string]int)
	var map2 = make(map[string]int)
	map3 := make(map[string]int)

	fmt.Println(map1)
	fmt.Println(map2)
	fmt.Println(map3)

	fmt.Println()

	map4 := map[string]int{}
	map4["apple"] = 25
	map4["banana"] = 40
	map4["orange"] = 33

	map5 := map[string]int{
		"apple":  15,
		"banana": 40,
		"orange": 23,
	}

	map6 := make(map[string]int, 10)
	map6["apple"] = 25
	map6["banana"] = 40
	map6["orange"] = 33

	fmt.Println(map4)
	fmt.Println(map5)
	fmt.Println(map6)
	fmt.Println(map6["apple"])
}
