package main


import (
	"fmt"
	"github.com/hangshiisi/stringutil"
	"time"
)

var c chan int

func ready(w string, sec int) {
	time.Sleep(time.Duration(sec) * time.Second)
	fmt.Println(w, "is ready!")

}

func readyTwo(w string, sec int) {
	time.Sleep(time.Duration(sec) * time.Second)
	fmt.Println(w, "is ready!")
	c <- 1
}

func mainFirst() {
	fmt.Println(stringutil.Reverse("!oG ,olleH"))
	go ready("Tea", 2)
	go ready("Coffee", 1)
	fmt.Println("I am waiting")
	time.Sleep(5 * time.Second)

}

func main() {
	fmt.Println(stringutil.Reverse("!oG ,olleH"))
	c = make(chan int)
	go readyTwo("Tea", 2)
	go readyTwo("Coffee", 1)
	fmt.Println("I am waiting, but not too long")
	//time.Sleep(5 * time.Second)
	<-c
	<-c

}
