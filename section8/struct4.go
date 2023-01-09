package main

import (
	"fmt"
)

type shoppingBasket struct{ cnt, price int }

func main() {

	bs1 := shoppingBasket{3, 5000}
	fmt.Println(bs1.purchase())

	bs1.rePurchaseP(7, 5000)
	fmt.Println(bs1.purchase())

	bs1.rePurchaseD(10, 0)
	fmt.Println(bs1.purchase())

}

func (b shoppingBasket) purchase() int {
	return b.cnt * b.price
}

// 원본 수정(참조형식)
func (b *shoppingBasket) rePurchaseP(cnt, price int) {
	b.cnt += cnt
	b.price += price
}

// 원본 수정 x( 값 전달 형식)
func (b shoppingBasket) rePurchaseD(cnt, price int) {
	b.cnt += cnt
	b.price += price
}
