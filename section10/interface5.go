package main

import "fmt"

type Dog struct {
	name   string
	weight int
}

type Cat struct {
	name   string
	weight int
}

func printValue(s interface{}) {
	fmt.Println("ex1 : ", s)
}

func main() {

	dog := Dog{"Poll", 1}
	cat := Cat{"marry", 2}

	printValue(dog)
	printValue(cat)
	printValue(12)
	printValue(13)
	printValue("aidj")

}
