package main

import "fmt"

func main() {
	var ptr *int

	var val int = 42

	ptr = &val

	fmt.Printf("ptr이 가리키는 값: %d\n", *ptr)

	*ptr = 100

	fmt.Printf("val의 새로운 값: %d\n", val)
}
