// Copyright © 2016 Hang Shi.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Policy Manager framework 
 
package main

import (
	"container/heap"
	"fmt"
	"os" 
//	"math/rand"
//	"time"
)

//=========================================================
//========================================================

type Request struct {
	fn func(w *Worker) int // The operation to perform.
	c  chan int   // The channel to return the result.
}

type Worker struct {
	requests chan Request // work to do (buffered channel)
	id       int //id of the worker 
	desc     string //description of the worker 
	pending  int // count of pending tasks
	index    int // index in the heap
}

type Pool []*Worker // used to build priority queue of workers 

type Balancer struct {
	pool Pool 		// priority queue for workers  
	done chan *Worker 	// channel for work requests
}

var (
	nWorker	int = 6		// number of workers 
	nJobs	int = 30	// how many job requests to start 
	completedJobs int = 0   // how many jobs have been completed 
	workChan	= make(chan Request)	// job request queue  
	allDoneChan	= make(chan int) 	// signal program bcompletion 	
)

func workFn(w *Worker) int {
	fmt.Println("Doing work inside function workFn by Worker ")
	fmt.Printf("\t   Details: id %d desc %s pending %d \n", 
                    w.id, w.desc, w.pending)  
	// time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	return 100
}

func furtherProcess(c int) int {
	return 200
}

// API to send requesters 
func requester(work chan<- Request) {
	for i := 0; i < nJobs ; i++ {
		// simulate the interval between job requests 
		// time.Sleep(time.Duration(rand.Intn(nWorker*2)) * time.Second)
		fmt.Printf("send Job request %d \n", i) 
		c := make(chan int)
		work <- Request{workFn, c} // send request
		result := <-c              // wait for answer
		close(c) 
		furtherProcess(result)
	}
	fmt.Printf("Exitting requester \n") 
}

func doTheWork(w *Worker, b *Balancer) {
	for req := range w.requests { // get Request from balancer
		req.c <- req.fn(w)    // call fn and send result
		fmt.Println("invoking worker functions") 
		b.done <- w           // we've finished this request
	}
	fmt.Printf("Exitting Worker \n") 
	
}

//Priority Queue Implementation 
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

func policyScheduler(b *Balancer, work <-chan Request) {
	loop: for {
		select {
		case req, ok := <- work: // received a Request...
			if !ok { 
				break loop
			} 
			b.dispatch(req) // ...so send it to a Worker
		case w, ok := <-b.done: // a worker has finished ...
			if !ok { 
				break loop
			} 
			b.completed(w) // ...so update its info
		}
	}
}

// Dispatch Request to worker based on various policy 
func (b *Balancer) dispatch(req Request) {
	// Grab the least loaded worker...
	w := heap.Pop(&b.pool).(*Worker)
	// ...send it the task.
	w.requests <- req
	// One more in its work queue.
	w.pending++
	fmt.Printf(" Get the work, assigned to worker %d pending %d \n", 
			w.id, w.pending) 
	// Put it into its place on the heap.
	heap.Push(&b.pool, w)
}

// Job is complete; update heap
func (b *Balancer) completed(w *Worker) {
	// One fewer in the queue.
	w.pending--
	fmt.Printf(" Completed the work, by worker %d pending %d \n", 
			w.id, w.pending) 
	// Remove it from heap.
	heap.Remove(&b.pool, w.index)

	// Put it into its place on the heap.
	heap.Push(&b.pool, w)

	completedJobs++
	fmt.Printf(" current Job completed is %d \n", completedJobs) 

        
	//if completedJobs == nJobs   { 
	//	fmt.Printf(" ALL Job completed already \n") 
 	//		allDoneChan <- 100 // all jobs completed 
	//} 

}

func runPolicyManager() {
	fmt.Println("Policy Manager Demo Started.\n")

	//worker id and strings 
	items := map[int]string{
		3:"banana", 2:"apple", 4:"pear",
                10:"kiwi", 11:"melon", 12:"orange", 
	}

	// Create a priority queue, put the items in it, and
	// establish the priority queue (heap) invariants.
	pq := make(Pool, len(items))
	i := 0
	for value, desc := range items {
		pq[i] = &Worker{
			id: value,
			desc:desc, 
			pending: 0,
			index:   i,
			requests: make(chan Request), 
		}
		i++
	}
	heap.Init(&pq)

	//create policy manager 
	pm := &Balancer{ 
		pool:pq, 
		done: make(chan *Worker), 
		}
 
	//start the worker goroutines 
	for i := 0; i < len(items); i++ { 
		go doTheWork(pq[i], pm)
	} 

	//start the policy scheduler goroutines 
	go policyScheduler(pm, workChan) 

	//start the requester goroutines 
	go requester(workChan)  

	fmt.Print("wait for the end signal \n") 
	//wait for the end signal 	
	<- allDoneChan 
	close(allDoneChan)
	close(workChan)
	//start the worker goroutines 
	for i := 0; i < len(items); i++ { 
		fmt.Printf("Close channel request %d \n", i) 
		close(pq[i].requests)
	} 
	close(pm.done) 
}

func testPQ() {
	fmt.Println("Hello World.\n")

	items := map[string]int{
		"banana": 3, "apple": 2, "pear": 4,
	}

	// Create a priority queue, put the items in it, and
	// establish the priority queue (heap) invariants.
	pq := make(Pool, len(items))
	i := 0
	for value, pending := range items {
		pq[i] = &Worker{
			id: i, 
			desc:      value,
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
		desc:       "Kiwi",
		id: 100, 
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
	go func() { 
		os.Stdin.Read(make([]byte,1)) 
		allDoneChan <- 100
	} () 	
	runPolicyManager()
} 



