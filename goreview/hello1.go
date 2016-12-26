package main

import "fmt"

func reviewFunction1() {

}

func rec(i int) {

	if i == 10 {
		return
	}

	rec(i + 1)
	fmt.Printf("%d ", i)
}

//named return arguments
//deferred codes, LIFO order
//defer function literals, need () there

func reviewDefer() (ret int) {
	defer func() {
		ret++
		ret++
	}()

	defer func() {
		ret++
		ret++
	}()

	return 0
}

//variadic parameters
//
func myvfunc(arg ...int) {
	for i, n := range arg {
		fmt.Printf("And the number is: %d %d\n", i, n)
	}
}

func myvfunc2(arg ...int) {
	fmt.Printf("First time output: \n\n")
	myvfunc(arg...)
	fmt.Printf("Second time output: \n\n")
	myvfunc(arg[:2]...)
}

//function as values

func reviewFuncValues() func() {

	a := func() {
		println("Hello Function Values")
	}
	a()
	return a
}

//func array return func values
func reviewFuncArray(i int) func() int {

	var xs = map[int]func() int{
		1: func() int { return 10 },
		2: func() int { return 20 },
		3: func() int { return 30 },
		/* ... */
	}

	return xs[i]
}

func main() {

	reviewFunction1()
	rec(0)
	k := reviewDefer()

	fmt.Printf("\nreturned values are %d \n", k)

	//myvfunc(1, 3, 5, 7, 9)
	//myvfunc2(2, 4, 6, 8, 10)

	fv := reviewFuncValues()
	fv()

	fv1 := reviewFuncArray(1)
	fmt.Printf("the value return is %d \n", fv1())

	fv2 := reviewFuncArray(2)
	fmt.Printf("the value return is %d \n", fv2())

}
