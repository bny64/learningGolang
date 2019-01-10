package main

import "fmt"

func main3() {
	const (
		Sunday = iota
		Monday
		Tuesday
		Wednesday
	)

	fmt.Println(Wednesday) //the usage of iota

	const (
		a = iota * 10
		b
		_
		c
	)

	fmt.Println(a, b, c)
}
