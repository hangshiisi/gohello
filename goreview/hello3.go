package main

import (
	"bytes"
	"fmt"
	"github.com/hangshiisi/gohello/goreview/even"
	"sync"
)

//review pointer basics
func reviewPointer1() {
	fmt.Printf("inside review Pointer ex1\n")
	var p *int
	fmt.Printf("%v \n", p)

	var i int
	p = &i
	fmt.Printf("%v %v\n", p, *p)

	*p = 8
	fmt.Printf("%v %v  value is %d \n", p, *p, i)

}

//two ways to allocate memory: new and make
//below is for new
func reviewPointerNew() {

	type SyncedBuffer struct {
		lock   sync.Mutex
		buffer bytes.Buffer
	}

	p := new(SyncedBuffer)
	var v SyncedBuffer

	fmt.Printf("%v \n\n %v\n", p, v)

}

//below is for make ( map, array, and channel only )
// new(T) returns *T pointing to a zeroed T
// make(T) returns an initialized T
func reviewPointerMake() {
	var p *[]int = new([]int)      //Allocates slice structure;rarely useful
	var v []int = make([]int, 100) //v refers to a new array of 100 ints

	fmt.Printf("p %v and v %v \n", p, v)

	var p1 *[]int = new([]int) //Unnecessarily complex
	*p1 = make([]int, 100, 100)
	fmt.Printf("p1 %v and *p1 %v \n", p1, *p1)

	v1 := make([]int, 100) //Idiomatic

	fmt.Printf("v1 %v\n", v1)

	/* make([]int, 10, 100) allocates an array of 100 ints and then creates
	 * a slice structure with length 10 and a capacity of 100 pointing at the
	 * first 10 elements of the array
	 */
}

func reviewStruct() {
	type T1 int
	type T2 int
	type T3 int

	type P4 struct {
		x, y int
	}

	type P5 struct {
		T1
		*T2
		T3
		x, y int
	}

	var pp5 P5

	fmt.Printf("inside review Struct %v \n", pp5)
}

func testEven() {
	i := 5
	fmt.Printf("Is %d even? \n %v \n", i, even.Even(i))
}

func Hello3Main() {
	//testEven()
	reviewPointer1()
	//reviewPointerNew()
	//reviewPointerMake()
	reviewStruct()
}
