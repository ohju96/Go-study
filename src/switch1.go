// 변수 1
package main

import "fmt"

func main() {
	a := 1
	switch {
	case a < 0:
		fmt.Println(a, "는 음수")
	case a == 0:
		fmt.Println(a, "는 0")
	case a > 0:
		fmt.Println(a, "는 양수")
	}

	switch b := 27; {
	case b < 0:
		fmt.Println(b, "는 음수")
	case b == 0:
		fmt.Println(b, "는 0")
	case b > 0:
		fmt.Println(b, "는 양수")
	}

	switch c := "tess"; c {
	case "go":
		fmt.Println("고")
	case "Java":
		fmt.Println("자바")
	default:
		fmt.Println("일치 없음")
	}

	switch c := "go"; c + "lang" {
	case "golang":
		fmt.Println("고랭")
	case "Java":
		fmt.Println("자바")
	default:
		fmt.Println("일치 없음")
	}

	switch i, j := 20, 30; {
	case i < j:
		fmt.Print("i는 j보다 작다")
	case i == j:
		fmt.Println("i와 j는 같다.")
	case i > j:
		fmt.Print("i는 j보다 크다")
	}
}
