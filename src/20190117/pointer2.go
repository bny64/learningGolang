package main

import "fmt"

//before main3
func main() {

	var num2 int
	fmt.Println(num2)
	fmt.Println(&num2)
	fmt.Println("----------")
	var num int = 1
	fmt.Println(num)
	fmt.Println(&num)
	//fmt.Println(*num)
	fmt.Println("----------")
	var numPtr *int = &num
	fmt.Println(&numPtr)
	fmt.Println(numPtr)
	fmt.Println(&num)
	fmt.Println(*numPtr)
}
