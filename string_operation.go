package main

import (
	"fmt"
	"strings"
)

func main() {

	var str1 string = "Golang"
	var str2 string = "World"

	fmt.Println(str1[0:2], str1[0])
	fmt.Println(str2[3:], str2[0])
	fmt.Println(str2[:4])
	fmt.Println(str2[1:3])

	fmt.Println("####")

	str01 := "Golang"
	str02 := "World"

	fmt.Println(str01 == str02)
	fmt.Println(str01 != str02)
	fmt.Println(str01 > str02)
	fmt.Println(str01 < str02)

	fmt.Println("####")

	str11 := "THe Go" +
		"programming" +
		"language"
	str12 := " Instructions" +
		"for"

	fmt.Println(str11 + str12)

	strSet := []string{}
	strSet = append(strSet, str1)
	strSet = append(strSet, str2)

	fmt.Println(strings.Join(strSet, `, `))

}
