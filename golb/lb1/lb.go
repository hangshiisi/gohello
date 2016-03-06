package main

import (
	"fmt"
	"time"
)

var NumWorkers int = 5

type Work struct {
	x, y, z int
}

func worker(in <-chan Work, out chan<- Work) {
	for w := range in {
		w.z = w.x * w.y
		// time.Sleep( w.z * time.Second)
		fmt.Println("Doing the work for %d ", w.x)

		time.Sleep(5 * time.Second)
		out <- w
	}
}

func sendLotsOfWork(in chan<- Work) {
	fmt.Println("Send, World")
	var w Work
	for i := 1; i <= 5; i++ {
		w.x = i
		w.y = i + 1
		w.z = 0
		in <- w
	}
	close(in)

}

func receiveLotsOfResults(out <-chan Work) {
	fmt.Println("Receive, World")
	for i := 1; i <= 5; i++ {
		w := <-out
		fmt.Println("Get Result %d ", w.z)
	}
}

func Run() {
	in, out := make(chan Work), make(chan Work)
	for i := 0; i < NumWorkers; i++ {
		go worker(in, out)
	}
	go sendLotsOfWork(in)
	receiveLotsOfResults(out)
	close(out)
}

func main() {
	Run()
}
