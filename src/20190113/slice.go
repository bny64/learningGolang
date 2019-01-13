package main

import "fmt"

//before 2
func main3() {
	//slice reference 타입
	var a []int = make([]int, 5)
	var b = make([]int, 5)
	c := make([]int, 5)

	_, _, _ = a, b, c

	var s = make([]int, 5, 10)
	fmt.Println(len(s), cap(s)) //length, capablity
}
