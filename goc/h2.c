#include <stdio.h> 
#include <stdlib.h> 
#include <time.h> 

//#include "_cgo_export.h"

int main2() 
{ 
    int i, j = 0, n, k; 
	clock_t start, end; 

	n = 100000000; 
	start = clock();

	for (i = 0; i < n; i++) { 
		//j++; 
		k = rand(); 
	}
	end = clock(); 

	printf("the time is %ld CLOCKS_PER_SEC %ld \n", 
			end - start, CLOCKS_PER_SEC); 
	printf("Hello World From C!\n"); 
	//HelloFromGo(); 
	return 0; 
}