package main

import "fmt"

//before 3
func main4() {
	a := []int{1, 2, 3}
	a = append(a, 4, 5, 6)
	b := []int{6, 7, 8}
	c := append(a, b...)
	fmt.Println(c)
	c[2] = 9
	fmt.Println(a)
	fmt.Println(c)

	d := make([]int, len(c))
	copy(d, c)
	d[4] = 11
	fmt.Println(c)
	fmt.Println(d)
}
