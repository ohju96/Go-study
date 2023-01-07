package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	var str1 string = "c:\\go_study\\src" // c:\go_study\src
	str2 := `c:\go_study\src`

	fmt.Println(str1)
	fmt.Println(str2)

	fmt.Println("####")

	var str3 string = "Hello, world!"
	var str4 string = "안녕하세요."
	var str5 string = "\ud55c\uae00"

	fmt.Println(str3)
	fmt.Println(str4)
	fmt.Println(str5)

	fmt.Println("####")

	// 바이트
	fmt.Println(len(str3))
	fmt.Println(len(str4))

	// 실제 길이
	fmt.Println(utf8.RuneCountInString(str3))
	fmt.Println(utf8.RuneCountInString(str4)) // 주로 선호
	fmt.Println(len([]rune(str4)))            // len으로 실제 길이 구하는 법

	fmt.Println("#####")

	var str01 string = "Golang"
	var str02 string = "World"
	var str03 string = "고프로그래밍"

	fmt.Println(str01[0], str01[1], str01[2], str01[3], str01[4], str01[5])
	fmt.Println(str02[0], str02[1], str02[3], str02[4])
	fmt.Println(str03[0], str03[1], str03[2], str03[3], str03[4], str03[5])

	fmt.Printf("%c, %c, %c, %c, %c, %c\n", str01[0], str01[1], str01[2], str01[3], str01[4], str01[5])
	fmt.Printf("%c, %c, %c, %c\n", str02[0], str02[1], str02[3], str02[4])
	fmt.Printf("%c, %c, %c, %c, %c, %c\n", str03[0], str03[1], str03[2], str03[3], str03[4], str03[5])

	conStr := []rune(str03)
	fmt.Printf("%c, %c, %c, %c, %c, %c\n", conStr[0], conStr[1], conStr[2], conStr[3], conStr[4], conStr[5])

	fmt.Println("#####")

	for i, char := range str01 {
		fmt.Printf("%c(%d)\t", char, i)
	}

	fmt.Println("####")

	for i, char := range str2 {
		fmt.Printf("%c(%d)\t", char, i)
	}
}
