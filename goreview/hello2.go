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

	// to extend slice, use append and copy
	s0 := []int{0, 0}
	fmt.Printf("slice append and copy %v \n", s0)

	s1 := append(s0, 2)
	fmt.Printf("s1 %v \n", s1)

	s2 := append(s1, 3, 5, 7)
	fmt.Printf("s2 %v \n", s2)

	s3 := append(s2, s0...)
	// s3 := append(s2, s0), cannot use s0 as type int in append

	fmt.Printf("s3 %v \n", s3)

	// slice copy
	var aa = [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	var ss = make([]int, 6)
	fmt.Printf("ss and aa original %v AND %v \n", aa, ss)
	n1 := copy(ss, aa[0:])
	fmt.Printf("ss now is %v copied %d \n", ss, n1)
	n2 := copy(ss, ss[2:])
	fmt.Printf("ss now is %v copied %d \n", ss, n2)

}

func reviewMap() {
	monthdays := map[string]int{
		"Jan": 31, "Feb": 28, "Mar": 31,
		"Apr": 30, "May": 31, "Jun": 30,
		"Jul": 31, "Aug": 31, "Sep": 30,
		"Oct": 31, "Nov": 30, "Dec": 31,
	}

	mds := make(map[string]int, 12)

	// mfd := copy(mds, monthdays...)
	// fmt.Printf("month days is %v copied \n", mds, mfd)
	// wrong, can't slice map

	alldays := 0
	i := 0
	for k, v := range monthdays {
		alldays += v
		mds[k] = v
		i++
	}
	fmt.Printf("all days total is %d \n", alldays)
	fmt.Printf("mds keys are %v \n", mds)
	fmt.Printf("original map is %v \n", monthdays)

}

func test() {
}

func Hello2Main() {

	reviewArray()
	reviewMap()

	//reviewSlice()
	//fallThrough(0)
	//fallThrough(1)
	//fallThrough(2)

}
