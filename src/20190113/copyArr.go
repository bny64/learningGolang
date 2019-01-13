package main

import "fmt"

//before 1
func main2() {
	a := [5]int{1, 2, 3, 4, 5}
	b := a
	//배열 복사는 새로운 배열을 생성
	a[1] = 10
	fmt.Println(a, b)
}
