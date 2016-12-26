package main

import "fmt"

func main() {
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

}
