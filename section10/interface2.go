package main

import "fmt"

func main() {

	dog1 := Dog{"poll", 19}

	var inter1 Behavior
	inter1 = dog1
	inter1.bite()

	fmt.Println()

	dog2 := Dog{"marry", 2}
	inter2 := Behavior(dog2)
	inter2.bite()

	fmt.Println()

	inters := []Behavior{dog1, dog2}
	for idx, _ := range inters {
		inters[idx].bite()
	}

	fmt.Println()

	for _, val := range inters {
		val.bite()
	}

}

type Dog struct {
	name   string
	weight int
}

func (d Dog) bite() {
	fmt.Println(d.name, "bites !")
}

type Behavior interface {
	bite()
}
