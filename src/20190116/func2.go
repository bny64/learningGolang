package main

import "fmt"

//before main1
func main2() {
	r := sumsum(2, 3, 4, 5, 6)
	fmt.Println(r)

	f := factorial(10)
	fmt.Println(f)
}

func sumsum(n ...int) int {
	total := 0
	for _, value := range n {
		total += value
	}

	return total
}

func factorial(n uint64) uint64 {
	if n == 0 {
		return 1
	}

	return n * factorial(n-1)
}
