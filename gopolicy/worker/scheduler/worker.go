package scheduler

import (
	"fmt"
	"math/rand"
	"time"
	// "net" //IP
)

// policy manager and policy agent communicates via REST APIs

// Policy Agent Message format among channels used by goroutinges
type PAMessage struct {
	code    int    // message types, type is go keyword!
	subtype int    // message subtypes
	dest    string // where to execute this message, IP
	nonce   int    // what is this message belongs to
	content string // for now, json string to host the rest
}

// multiple channels:
//	* inWorkChan, to get the work sent from PM (policy manager)
// 	* outResultChan, to send the result to PM (policy manager)
//	* (future) reportChan, to send reports to PM (policy manager)
//		(may consider combining with outResultChan
// each channel will have a dedicated listener which behaves in a non-blocking
//	fashion and will create goroutine for incoming work.
// multiple goroutines will be created to achieve concurrency.
// Some additional channels will be used for synchronization.
// no blocking or locks will be done in this.
//      * system flow for requests
//              work requests will be sent from policy manager and REST
//		server will receive these requests via JSON format and create
// 		PAMessages and send them to the request listener goroutine
// 		via inWorkChan. Listener goroutine will then create a goroutine
//		to do the work for each request; when the work is done, that
//		goroutine will post the reply to outResultChan.
//      * system flow for replies
// 		the replies will correspond to previous requests. the reply
//		listener goroutine will get the replies via outResultChan. It
//		will then create
//	* system flow for reports
//		the health report, membership report, anomaly report, etc.
//		for each function, a goroutine can be created to do its job
//		. Results from these functions will be posted to the
// 		reportChan and aggregated to policy manager.
// the message format will need further investigation, more details will be
// 	added and different channels may use different message formats.
// Note this message format is internal and different from external i
//	communication.
//
// goroutine summary:
//	* one goroutine as listener for each channel:
// 		* inWorkListener, to handle job requests
// 		* outResultListener, to handle job replies
// 		* reportLister (future), to handle reports
//      * for each work sent to the listener, one goroutine will be created
// 	* goroutines created by REST/Web servers
// 	* shell support (future)
//
var inWorkChan = make(chan PAMessage)
var outResultChan = make(chan PAMessage)
var reportChan = make(chan PAMessage)

// in policy agent, currently we support one worker
var NumWorkers int = 1

// indicate the completion of this application
var doneChan = make(chan int)

func inWorkListener(myWorkerID int, in <-chan PAMessage,
	out chan<- PAMessage) {
	for w := range in {
		//simulating the load
		k := rand.Intn(NumWorkers * 2)
		fmt.Printf("I'm %d job %d sleep %d\n", myWorkerID,
			w.code, k)
		time.Sleep(time.Duration(k) * time.Second)
		w.nonce = w.code * w.subtype
		out <- w
	}
	close(out)
}

// this function simulates the work from clients
func testWorkRequester(in chan<- PAMessage) {
	fmt.Println("Start Goroutine for Job Requests")
	for i := 1; i <= 5; i++ {
		in <- PAMessage{code: i, subtype: i + 1}
	}
	close(in)
}

func outResultListener(out <-chan PAMessage) {
	fmt.Println("Start Goroutine for Result Retrieval")
	for w := range out {
		fmt.Printf("Get Result %d for request No. %d \n",
			w.nonce, w.subtype)
	}
	doneChan <- 100
}

func RunPolicyAgent() {

	for i := 0; i < NumWorkers; i++ {
		go inWorkListener(i, inWorkChan, outResultChan)
	}

	go testWorkRequester(inWorkChan)

	go outResultListener(outResultChan)

	<-doneChan

}

func ownMain() {
	RunPolicyAgent()
}
