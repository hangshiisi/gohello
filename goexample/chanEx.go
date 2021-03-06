package main

import (
	"fmt"
	"time"
)

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

func showerMain() {
	ch := make(chan int)
	go shower(ch)
	for i := 1; i < 10; i++ {
		ch <- i
	}
}

func shower(ch chan int) {
	for {
		fmt.Printf("%d \n", <-ch)
	}
}

func dup3(in chan int) (chan int, chan int, chan int) {

	a, b, c := make(chan int, 2), make(chan int, 2), make(chan int, 2)
	go func() {
		for {
			x := <-in
			a <- x
			b <- x
			c <- x
		}
	}()
	return a, b, c
}

func fib() chan int {
	x := make(chan int, 2)
	a, b, out := dup3(x)
	go func() {
		x <- 0
		x <- 1
		<-a
		for {
			x <- <-a + <-b
		}
	}()
	return out
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	x := fib()
	for i := 0; i < 10; i++ {
		fmt.Println(<-x)
	}

	go say("world")
	say("hello")

}
