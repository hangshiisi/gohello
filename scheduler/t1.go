package main
import "fmt"
import "time"
import "runtime" 


func main()  { 

	var x int
	threads := runtime.GOMAXPROCS(0) 
	for i := 0; i < threads; i++ { 
		go func() { 
			for { x++ } 
		} () 
	} 
	time.Sleep(time.Second * 2) 
	fmt.Println("x =  ", x) 
} 


