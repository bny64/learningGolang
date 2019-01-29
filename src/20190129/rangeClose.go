package main

import "fmt"

func main() {
	/* c := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c)
	}()

	for i := range c {
		fmt.Println(i)
	} */

	c := make(chan int, 1)

	go func() {
		c <- 1
	}()

	a, ok := <-c
	fmt.Println(a, ok)

	close(c)
	a, ok = <-c
	fmt.Println(a, ok)
}
