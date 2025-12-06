package main

import "fmt"

func main() {
	var nilValue []string

	fmt.Printf("Is Nil?: %t\n", nilValue == nil)

	nilValue = []string{"hello"}

	fmt.Printf("Is Nil?: %t\n", nilValue == nil)
}
