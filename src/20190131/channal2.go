package main

import "fmt"

func sum(a, b int) <-chan int {
	out := make(chan int)

	go func() {
		out <- a + b
	}()

	return out
}

func main2() {
	c := sum(1, 2) //chan type
	fmt.Println(<-c)
}
