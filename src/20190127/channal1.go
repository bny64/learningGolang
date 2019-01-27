package main

import (
	"fmt"
	"time"
)

func sum(a int, b int, c chan int) {

	time.Sleep(2000 * time.Millisecond)
	c <- a + b
}

func main5() {
	c := make(chan int)
	go sum(1, 2, c)

	n := <-c
	fmt.Println(n)
}
