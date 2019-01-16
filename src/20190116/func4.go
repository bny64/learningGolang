package main

import "fmt"

func calc() func(x int) int {
	a, b := 3, 5
	return func(x int) int {
		return a*x + b
	}
}

func main4() {
	f := calc()

	fmt.Println(f(1))
	fmt.Println(f(3))
}
