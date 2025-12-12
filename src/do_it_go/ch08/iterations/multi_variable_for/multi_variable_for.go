package main

import "fmt"

func main() {
	for i, j := 0, 5; i < j; i, j = i+1, j-1 {
		fmt.Printf("i = %d, j = %d\n", i, j)
	}

	var iter int = 100
	for i := 0; i < iter; i++ {
		fmt.Printf("%d ", i+iter)
	}
}
