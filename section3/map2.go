package main

import "fmt"

func main() {

	map1 := map[string]string{
		"daum":   "http://daum.net",
		"naver":  "http://naver.com",
		"google": "http://google.com",
	}

	fmt.Println(map1["google"])
	fmt.Println(map1["daum"])

	fmt.Println()

	for k, v := range map1 {
		fmt.Println(k, v)
	}

	fmt.Println()

	for _, v := range map1 {
		fmt.Println(v)
	}

	fmt.Println()

	map2 := map[string]string{
		"daum":   "http://daum.net",
		"naver":  "http://naver.com",
		"google": "http://google.com",
	}

	map2["home"] = "http://test.com"
	fmt.Println("home : ", map2)

	fmt.Println()

	map2["home"] = "http://test2com"
	fmt.Println("home : ", map2)

	fmt.Println()

	delete(map1, "home")
	fmt.Println(map1)

}
