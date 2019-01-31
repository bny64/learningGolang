package main

import "fmt"

func num3(a, b int) <-chan int {
	out := make(chan int)
	go func() {
		out <- a
		out <- b
		close(out)
	}()

	return out
}

func sum3(c <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		r := 0
		for i := range c {
			r = r + i
		}
		out <- r
	}()
	return out
}

func main3() {
	c := num3(1, 2)
	out := sum3(c)

	fmt.Println(<-out)
}
