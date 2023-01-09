package main

import "fmt"

func main() {

	car1 := struct{ name, color string }{"520d", "red"}
	fmt.Println(car1)
	fmt.Printf("%#v\n", car1)

	cars := []struct{ name, color string }{
		{"520d", "red"},
		{"530i", "white"},
	}

	for _, c := range cars {
		fmt.Printf("(%s, %s) ----- (%#v)\n", c.name, c.color, c)
	}
}
