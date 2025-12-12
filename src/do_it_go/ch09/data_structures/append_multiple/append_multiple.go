package main

import "fmt"

func main() {
	var sliceValue []int

	sliceValue = append(sliceValue, 1, 2, 3, 4, 5)

	fmt.Printf("sliceValue len: %d, cap: %d, value: %v\n", len(sliceValue), cap(sliceValue), sliceValue)
}
