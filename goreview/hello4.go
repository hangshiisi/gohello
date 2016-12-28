package main

import (
	"fmt"
	"reflect"
)

type S struct {
	i int
}

func (p *S) Get() int  { return p.i }
func (p *S) Put(v int) { p.i = v }

func reviewInterface1() {
	fmt.Printf("Inside review interface\n")
	var k S
	fmt.Printf("S value is %d \n", k.Get())

	k.Put(10)
	fmt.Printf("S value is %d \n", k.Get())

}

type I interface {
	Get() int
	Put(int)
}

func ff(p I) {

	// another way to find if a type implements an interface is
	// to use "comma, ok" form during runtime
	if t, ok := p.(I); ok {
		// something implements the interface I
		// t is the type it has
		fmt.Printf("t has the type %v \n", t)
	}

	// can use switch t := p.(type) to find the type of p
	// using (type) outside a switch is illegal

	fmt.Printf("S value 1 is %d \n", p.Get())
	p.Put(100)
	fmt.Printf("S value 2 is %d \n", p.Get())

}

type R struct{ i int }

func (p *R) Get() int  { return p.i }
func (p *R) Put(v int) { p.i = v }

func reviewInterface2() {
	var k S

	k.Put(390)

	var dd interface{}
	dd = &k
	if t, ok := dd.(I); ok {
		// something implements the interface I
		// t is the type it has
		fmt.Printf("t now has the type %v \n", t)
	}

	ff(&k)

	var t R
	ff(&t)

}

type Foo int

func (self Foo) Emit() {
	fmt.Printf("%v\n", self)
}

type Emitter interface {
	Emit()
}

func reviewInterface3() {
	var kf Foo

	fmt.Printf("inside review interface 3\n\n")
	kf = 321
	kf.Emit()

	ki := &kf
	ki.Emit()

}

//generic function example
//sort example
type Sorter interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

type Xi []int
type Xs []string

func (p Xi) Len() int               { return len(p) }
func (p Xi) Less(i int, j int) bool { return p[j] < p[i] }
func (p Xi) Swap(i int, j int)      { p[i], p[j] = p[j], p[i] }

func (p Xs) Len() int               { return len(p) }
func (p Xs) Less(i int, j int) bool { return p[j] < p[i] }
func (p Xs) Swap(i int, j int)      { p[i], p[j] = p[j], p[i] }

func BubbleSort(x Sorter) {
	for i := 0; i < x.Len()-1; i++ {
		for j := i + 1; j < x.Len(); j++ {
			if x.Less(i, j) {
				x.Swap(i, j)
			}
		}
	}

}

func reviewInterface4() {
	ints := Xi{44, 67, 3, 17, 89, 10, 73, 9, 14, 8}
	strings := Xs{"nut", "ape", "elephant", "zoo", "go"}
	BubbleSort(ints)
	fmt.Printf("%v\n", ints)
	BubbleSort(strings)
	fmt.Printf("%v\n", strings)
}

type Person struct {
	// need to make the field capital case
	Name string "namestr" //tag
	Age  int
}

func showTag(i interface{}) {
	switch t := reflect.TypeOf(i); t.Kind() {
	case reflect.Ptr:
		tag := t.Elem().Field(0).Tag
		fmt.Printf("value of Tag is %v \n", tag)
	}
}

func showInfo(i interface{}) {
	switch i.(type) {
	case *Person:
		t := reflect.TypeOf(i)
		v := reflect.ValueOf(i)
		tag := t.Elem().Field(0).Tag
		name := v.Elem().Field(0).String()
		fmt.Printf("tag name info %v and %v \n", tag, name)
	}

}

/*
 * below function does not compile

func SetPrivate(i interface{}) {
	switch i.(type) {
	case *Person:
		r := reflect.ValueOf(i)
		r.Elem(0).Field(0).SetString("Albert Einstein")
	}
}
*/

func SetPublic(i interface{}) {
	switch i.(type) {
	case *Person:
		r := reflect.ValueOf(i)
		r.Elem().Field(0).SetString("Albert Einstein")
	}
}

func reviewInterface5() {
	var p1 Person
	SetPublic(&p1)

	showTag(&p1)
	showInfo(&p1)

}

func Hello4Main() {

	reviewInterface1()
	//reviewInterface2()
	//reviewInterface3()
	//reviewInterface4()
	reviewInterface5()

}
