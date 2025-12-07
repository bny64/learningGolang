package main

import "fmt"

func main() {
	increase := 1
	increase++
	fmt.Println(increase)

	var decrease = 1
	decrease--
	fmt.Println(decrease)

	name := "inho"
	age := 20

	fmt.Println(name + " is " + fmt.Sprint(age) + "years old")
}
