package main

import (
	"fmt"
)

type Rectangle1 struct {
	width, height int
	name          string
}

func main1() {
	rect1 := new(Rectangle1)
	rect2 := Rectangle1{}
	fmt.Println(rect1)
	fmt.Println(rect2)

	rect1.name = "hello"
	rect2.name = "power"
	fmt.Println(*rect1)
	fmt.Println(rect2)
}
