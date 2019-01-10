package main

import (
	. "fmt"       //global package. we have to use this method only special period
	run "runtime" //we can make special nickname at package and use
	_ "unsafe"    //when we do not use package, we just put under bar
)

func main4() {
	a := 2
	b := 3
	Println(a & b) // bit operator
	//more information about that we don't use normaly is in page 60, 61.
	Println(run.NumCPU())
}
