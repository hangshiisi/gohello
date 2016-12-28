package main

import (
	"fmt"
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
	fmt.Printf("S value 1 is %d \n", p.Get())
	p.Put(100)
	fmt.Printf("S value 2 is %d \n", p.Get())
}

func reviewInterface2() {
	var k S
	ff(&k)

}

func Hello4Main() {

	reviewInterface1()
	reviewInterface2()

}
