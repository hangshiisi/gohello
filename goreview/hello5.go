package main

import (
	"fmt"
	"time"
)

//channels and goroutines
/*
  what are goroutines
  They’re called goroutines because the existing terms — threads,
  coroutines, pro- cesses, and so on — convey inaccurate connotations.
  A goroutine has a simple model: it is a function executing in
  parallel with other goroutines in the same address space. It is
  lightweight, costing little more than the allocation of stack
  space. And the stacks start small, so they are cheap, and grow by
  allocating (and freeing) heap storage as required.
*/

var c1, c2 chan int

func ready(c chan int, w string, sec int) {

	time.Sleep(time.Duration(sec) * time.Second)
	fmt.Println(w, " is ready")
	c <- 1
}

func reviewChannel1() {
	fmt.Printf("inside review Channel1 \n")
	c1 = make(chan int)
	go ready(c1, "Tead", 2)
	go ready(c1, "Coffee", 2)
	fmt.Println("I'm waiting, but not too long")
	<-c1
	<-c1
}

func reviewChannel2() {
	fmt.Printf("inside review Channel2 \n")

	c2 = make(chan int)
	go ready(c2, "Tead", 2)
	go ready(c2, "Coffee", 2)
	fmt.Println("I'm waiting, but not too long")
	i := 0
L:
	for {
		select {
		case <-c2:
			i++
			if i > 1 {
				break L
			}
		}
	}
}

//buffered channel and unbuffered channel
//closing channel
/*
The following code will check if a channel is closed.
		x, ok = <-ch
Where ok is set to true the channel is not closed and we’ve read something. Otherwise
ok is set to false. In that case the channel was closed.
*/

func Hello5Main() {
	reviewChannel1()
	reviewChannel2()
}
