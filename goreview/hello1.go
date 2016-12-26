package main

import "fmt"

func review1() {
	var a, b = 6, 7
	c1, d := 8, 9

	var s string = "test string "

	fmt.Printf("Hello, World %d %d \n", a, b)
	fmt.Printf("Hello, World %d %d \n", c1, d)

	fmt.Printf("Hello, String %s \n", s)

	s = "test2 string"

	fmt.Printf("Hello, String %s \n", s)

	c := []rune(s)
	c[0] = 'D'
	c[1] = 'T'
	s2 := string(c)
	fmt.Printf("Hello, String %s \n", s2)

	for i := 1; i < 8; i++ {
		fmt.Printf("Hello, Value %d \n", i)
	}

}

func review2() {
	list := []string{"a", "b", "c", "d", "e", "f"}

	for k, v := range list {
		fmt.Printf("k, v is: %d %s \n", k, v)

	}

	for pos, char := range "xyz" {
		fmt.Printf("character '%c' starts at byte position %d \n",
			char, pos)

	}
}

func main() {

	review2()
}
