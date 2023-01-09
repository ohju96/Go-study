package main

import "fmt"

func main() {

	dog := Dog{"Poll", 1}
	cat := Cat{"marry", 2}

	act(dog)
	act(cat)

	// 덕 타이핑 : 구조체 및 변수의 값이나 타입은 상관하지 않고 오로지 구현된 메소드로만 판단하는 방식

}

type Dog struct {
	name   string
	weight int
}

type Cat struct {
	name   string
	weight int
}

type Behavior interface {
	bite()
	sounds()
	run()
}

func (d Dog) bite() {
	fmt.Println(d.name, " Dog bites !")
}

func (d Dog) sounds() {
	fmt.Println(d.name, " Dog barks !")
}

func (d Dog) run() {
	fmt.Println(d.name, " Dog running !")
}

func (d Cat) bite() {
	fmt.Println(d.name, " Cat 할퀴다 !")
}

func (d Cat) sounds() {
	fmt.Println(d.name, " Cat barks !")
}

func (d Cat) run() {
	fmt.Println(d.name, " Cat running !")
}

func act(animal Behavior) {
	animal.bite()
	animal.sounds()
	animal.run()

}
