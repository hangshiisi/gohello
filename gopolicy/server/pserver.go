// Copyright © 2016 Hang Shi.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

//
//!+

// Chat is a server that lets clients chat with each other.
package main

import (
	"container/heap"
	"fmt"
	"math/rand"
	"time"
)

//=========================================================
//========================================================

type Request struct {
	fn func() int // The operation to perform.
	c  chan int   // The channel to return the result.
}

var (
	nWorker	int64 = 4
	nJobs	int = 10
	workChan       = make(chan Request)
)

func workFn() int {
	fmt.Println("inside function workFn \n")
	return 100
}

func furtherProcess(c int) int {
	fmt.Println("inside function furtherProcess \n")
	return 200
}

func requester(work chan<- Request) {
	for i := 1; i < nJobs ; i++ {
		// simulate the load
		time.Sleep(time.Duration(rand.Int63n(nWorker*2)) * time.Second)
		c := make(chan int)
		work <- Request{workFn, c} // send request
		result := <-c              // wait for answer
		close(c) 
		furtherProcess(result)
	}
}

type Worker struct {
	requests chan Request // work to do (buffered channel)
	id       string
	pending  int // count of pending tasks
	index    int // index in the heap
}

func (w *Worker) work(done chan *Worker) {
	for req := range w.requests { // get Request from balancer
		req.c <- req.fn()   // call fn and send result
		done <- w           // we've finished this request
	}
}

type Pool []*Worker

type Balancer struct {
	pool Pool
	done chan *Worker
}

func (b *Balancer) balance(work chan Request) {
	for {
		select {
		case req := <-work: // received a Request...
			b.dispatch(req) // ...so send it to a Worker
		case w := <-b.done: // a worker has finished ...
			b.completed(w) // ...so update its info
		}
	}
}

func (p Pool) Less(i, j int) bool {
	return p[i].pending < p[j].pending
}

func (h Pool) Len() int { return len(h) }

func (h Pool) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].index = i
	h[j].index = j
}

func (h *Pool) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	n := len(*h)
	item := x.(*Worker)
	item.index = n
	*h = append(*h, item)

}

func (h *Pool) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	x.index = -1
	*h = old[0 : n-1]
	return x
}

func (h *Pool) update(item *Worker, pending int) {
	item.pending = pending
	heap.Fix(h, item.index)
}

// Send Request to worker
func (b *Balancer) dispatch(req Request) {
	// Grab the least loaded worker...
	w := heap.Pop(&b.pool).(*Worker)
	// ...send it the task.
	w.requests <- req
	// One more in its work queue.
	w.pending++
	// Put it into its place on the heap.
	heap.Push(&b.pool, w)
}

// Job is complete; update heap
func (b *Balancer) completed(w *Worker) {
	// One fewer in the queue.
	w.pending--
	// Remove it from heap.
	heap.Remove(&b.pool, w.index)
	// Put it into its place on the heap.
	heap.Push(&b.pool, w)
}

func runPolicyManager() {
	fmt.Println("Hello World.\n")

	// Some items and their priorities.
	//items := map[chan Request]int{
	// 		make(chan Request): 3, make(chan Request): 2,
	//		make(chan Request): 4,
	//	}
	items := map[string]int{
		"banana": 3, "apple": 2, "pear": 4,
	}

	// Create a priority queue, put the items in it, and
	// establish the priority queue (heap) invariants.
	pq := make(Pool, len(items))
	i := 0
	for value, pending := range items {
		pq[i] = &Worker{
			id:      value,
			pending: pending,
			index:   i,
		}
		i++
	}
	heap.Init(&pq)

	// Insert a new item and then modify its priority.
	item := &Worker{
		requests: make(chan Request),
		pending:  1,
		id:       "Kiwi",
	}
	heap.Push(&pq, item)
	pq.update(item, 5)

	// Take the items out; they arrive in decreasing priority order.
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Worker)
		fmt.Printf("%.2d:%s \n", item.pending, item.id)
	}
}

func testPQ() {
	fmt.Println("Hello World.\n")

	// Some items and their priorities.
	//items := map[chan Request]int{
	// 		make(chan Request): 3, make(chan Request): 2,
	//		make(chan Request): 4,
	//	}
	items := map[string]int{
		"banana": 3, "apple": 2, "pear": 4,
	}

	// Create a priority queue, put the items in it, and
	// establish the priority queue (heap) invariants.
	pq := make(Pool, len(items))
	i := 0
	for value, pending := range items {
		pq[i] = &Worker{
			id:      value,
			pending: pending,
			index:   i,
		}
		i++
	}
	heap.Init(&pq)

	// Insert a new item and then modify its priority.
	item := &Worker{
		requests: make(chan Request),
		pending:  1,
		id:       "Kiwi",
	}
	heap.Push(&pq, item)
	pq.update(item, 5)

	// Take the items out; they arrive in decreasing priority order.
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Worker)
		fmt.Printf("%.2d:%s \n", item.pending, item.id)
	}

}

//=========================================================
//========================================================
//=========================================================

func main() { 
	fmt.Println("Hello World")
	// testPQ()
} 



