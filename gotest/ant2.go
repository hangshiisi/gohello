package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Message struct {
	str  string
	wait chan bool
}

func boring(msg string, wait chan bool) <-chan Message {
	// Returns receive-only channel of strings.
	c := make(chan Message)
	go func() { // We launch the goroutine from inside the function.
		for i := 0; ; i++ {
			c <- Message{fmt.Sprintf("%s %d", msg, i), wait}
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
			<-wait
		}
	}()
	return c // Return the channel to the caller.
}

func fanIn(input1, input2 <-chan Message) <-chan Message {
	c := make(chan Message)
	go func() {
		for {
			select {
			case s := <-input1:
				c <- s
			case s := <-input2:
				c <- s
			}
		}

	}()
	return c
}

func testMain() {
	waitForIt := make(chan bool) // Shared between all messages.
	c := fanIn(boring("Joe", waitForIt), boring("Ann", waitForIt))

	for i := 0; i < 5; i++ {
		msg1 := <-c
		fmt.Println(msg1.str)
		msg2 := <-c
		fmt.Println(msg2.str)
		msg1.wait <- true
		msg2.wait <- true
	}

}

func main() {
	waitForIt := make(chan bool) // Shared between all messages.
	c := boring("Joe", waitForIt)
	timeout := time.After(5 * time.Second)
	for {
		select {
		case s := <-c:
			fmt.Println(s.str)
			s.wait <- true
		case <-timeout:
			fmt.Println("You talk too much.")
			return
		}
	}

}
