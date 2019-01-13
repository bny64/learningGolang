package main

import "fmt"

func main1() {
	var a [5]int = [5]int{1, 2, 3, 4, 5}
	var b = [5]int{1, 2, 3, 4, 5}
	c := [5]int{1, 2, 3, 4, 5}
	fmt.Println(a, b, c)

	d := [...]string{"a", "b", "c"}
	fmt.Println(d)

	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}

	for _, value := range d {
		fmt.Println(value)
	}
}
