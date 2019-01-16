package main

import "fmt"

func main3() {
	var hello = sum
	world := hello
	fmt.Println(world(4, 5))

	f := []func(int, int) int{mainSum2, mainDiff2}

	fmt.Println(f[0](1, 2))
	fmt.Println(f[1](3, 4))

	f2 := map[string]func(int, int) int{
		"sum":  mainSum2,
		"diff": mainDiff2,
	}

	fmt.Println(f2["sum"](4, 3))
	fmt.Println(f2["diff"](6, 7))
}

func mainSum(a int, b int) int {
	return a + b
}

func mainSum2(a int, b int) int {
	return a + b
}

func mainDiff2(a int, b int) int {
	return a - b
}
