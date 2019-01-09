package main

import (
	"fmt"
	"math"
)

func main2() {
	/* var num1 int = 32
	var num2 int = -15
	var num3 int = 0723
	var num4 int = 0x32fa2c75

	fmt.Println(num1, num2, num3, num4)

	var num5 float32 = 1e7

	fmt.Println(num5) */

	var a float64 = 10.0

	for i := 0; i < 10; i++ {
		a = a - 0.1
	}

	fmt.Println(a)

	const epsilon = 1e-14

	fmt.Println(math.Abs(a-9.0) <= epsilon)
}
