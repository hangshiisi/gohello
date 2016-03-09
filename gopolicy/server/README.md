
1. workChan: global variable to enqueu/dequeue work messages 
	pass it via parameters 
2. request specific channel c is passed as part of the request 
	in order to return the results from the request processing
3. Each worker will have a channel to enqueue work requests 
	and keep track of the pending job
	when the job is finished, the result is passed to done channel 

4. doneChan:   
	a channel inside Load Balancer to enqueue/dequeue done messages 

5. allDoneChan: 
	to indicate all work has been completed 

Flow
	create worker states/priority queue
	create policy manager 
	start the worker goroutine (N) 
	start policy manager goroutine (1) 
	start the API request goroutine (1 for now) 
	wait for all done sigal 

