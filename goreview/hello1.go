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

func fallThrough(i int) {

	switch i {
	case 0:
		fallthrough
	case 1:
		fmt.Printf(" Falling through action \n")
	}
}

func reviewArray() {
	arr := [...]int{1, 2, 3}
	arr[0] = 42
	arr[1] = 32

	fmt.Printf("array elements %d %d %d \n", arr[0], arr[1], arr[2])

}

func reviewSlice() {
	sl := make([]int, 10)
	fmt.Printf("slice length %d and capacity %d \n", len(sl), cap(sl))

	var array [100]int
	slice := array[0:99]

	slice[98] = 'a'
	// slice[99] = 'a' 'index out of range'

	a := [...]int{1, 2, 3, 4, 5} //define an array with 5 elements
	fmt.Printf("a  %v \n", a)

	a1 := a[2:4] // slice with elements 3, 4
	fmt.Printf("a1 %v \n", a1)

	a2 := a[1:5] // slice with elements 2, 3, 4, 5
	fmt.Printf("a2 %v \n", a2)

	a3 := a[:] // same as a[0:len(a)], slice with elements 1, 2, 3, 4, 5
	fmt.Printf("a3 %v \n", a3)

	a4 := a[:4] // same as a[0:4], slice with elements 1, 2, 3, 4
	fmt.Printf("a4 %v \n", a4)

	a5 := a2[:] // same as a2[0:len(a2)]
	fmt.Printf("a5 %v \n", a5)

	a6 := a[2:4:5] // slice with elements from index
	fmt.Printf("a6 %v \n", a6)

}

func main() {

	reviewArray()
	reviewSlice()

	//fallThrough(0)
	//fallThrough(1)
	//fallThrough(2)

}
