package main

/*
#include <stdio.h>

extern int sum(int a, int b);

static inline void CExample(){
	int r = sum(1,2);
	printf("%d\n", r);
}
*/
import "C"

func sum(a, b C.int) C.int {
	return a + b
}

func main() {
	C.CExample()
}
