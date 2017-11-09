#include <stdio.h> 
#include <stdlib.h> 
#include <time.h> 


int main() 
{ 
	int i, j = 0, n, k; 
	time_t t1, t2; 


	clock_t start, end; 


	n = 10000000; 

	srand((unsigned)time(&t1)); 
	printf("hello world\n"); 
	
	t1 = time(NULL); 
	start = clock();

	for (i = 0; i < n; i++) { 
		j++; 
		k = rand(); 
	}
	end = clock(); 

	t2 = time(NULL); 

	printf("the time is %ld \n", t2 - t1); 
	printf("the time is %ld CLOCKS_PER_SEC %ld \n", 
			end - start, CLOCKS_PER_SEC); 
	printf("the final j is %d \n", j); 
	return 0; 
}