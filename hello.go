
package main 

import "fmt"
import "time"
import "math/rand"

// #include <stdio.h>
// #include <stdlib.h>
import "C"

func Random() int {
	return int(C.random())
}

func main() {

	t0 := time.Now()
	fmt.Println("Hello, ??")
	
	t1 := time.Now()
	dur := t1.Sub(t0)
	fmt.Printf(" t0 %v \n", t0)
	fmt.Printf(" t1 %v dur %v \n", t1, dur)
	
	t2 := time.Now()
	x := 100
	for i := 1; i < 1000; i++ {
		x = x + 1
	}
	t3 := time.Now()
	dur2 := t3.Sub(t2)
	c_rand := Random()
	fmt.Printf(" x %v, c_rand: %v, dur2: %v \n", x, c_rand, dur2)
	
	t4 := time.Now()
	for i := 1; i < 1000; i++ {
		//x = x + 1
		c_rand = Random()
		// c_rand = int(C.random())
	}
	t5 := time.Now()
	dur3 := t5.Sub(t4)
	fmt.Printf(" x %v, c_rand: %v, dur3: %v \n", x, c_rand, dur3)


	r := rand.New(rand.NewSource(99))
	t6 := time.Now()
	for i := 1; i < 1000; i++ {
		//x = x + 1
		c_rand = r.Intn(10000000000)
	}

	t7 := time.Now()
	dur4 := t7.Sub(t6)
	fmt.Printf(" x %v, c_rand: %v, dur4: %v \n", x, c_rand, dur4)
}
