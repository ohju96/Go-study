package main

import "fmt"

func main() {

	dog := Dog{"Poll", 1}
	cat := Cat{"marry", 2}

	act(dog)
	act(cat)

}

type Dog struct {
	name   string
	weight int
}

type Cat struct {
	name   string
	weight int
}

func (d Dog) run() {
	fmt.Println(d.name, " Dog running !")
}

func (d Cat) run() {
	fmt.Println(d.name, " Cat running !")
}

func act(animal interface{ run() }) {
	animal.run()
}
