package main

import (
	"fmt"
	"time"
	"math/rand"
)

var NumWorkers int = 5
var doneChan = make(chan int) 

type Work struct {
	x, y, z, k int
}

func worker(in <-chan Work, out chan<- Work) {
	for w := range in {
                //simulating the load 
		w.k = rand.Intn(NumWorkers*2) 
		fmt.Printf("Doing the work for %d sleep %d\n", w.x, w.k)
		time.Sleep(time.Duration(w.k) * time.Second)
		w.z = w.x * w.y
		out <- w
	}
}

func sendLotsOfWork(in chan<- Work) {
	fmt.Println("Start Goroutine for Job Requests")
	for i := 1; i <= 5; i++ {
		in <- Work{x:i, y:i+1}  
	}
	close(in)

}

func receiveLotsOfResults(out <-chan Work) {
	fmt.Println("Start Goroutine for Result Retrieval")
	for i := 1; i <= NumWorkers; i++ {
		w := <-out
		fmt.Printf("Get Result %d for request No. %d \n", w.z, w.x)
	}
        doneChan <- 100 
}

func Run() {
	in, out := make(chan Work), make(chan Work)
	for i := 0; i < NumWorkers; i++ {
		go worker(in, out)
	}
	go sendLotsOfWork(in)
	go receiveLotsOfResults(out)
        <- doneChan 
	close(out)
}

func main() {
	Run()
}
