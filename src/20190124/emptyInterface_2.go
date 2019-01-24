package main

import (
	"fmt"
	"strconv"
)

type Person2 struct {
	name string
	age  int
}

func formatString2(arg interface{}) string {
	switch arg.(type) {
	case Person2:
		p := arg.(Person2)
		return p.name + " " + strconv.Itoa(p.age) + "normal"
	case *Person2:
		p := arg.(*Person2)
		return p.name + " " + strconv.Itoa(p.age) + "pointer"
	default:
		return "Error"
	}
}

func main() {
	fmt.Println(formatString2(Person2{"Maria", 20}))
	fmt.Println(formatString2(&Person2{"John", 12}))

	var andrew = new(Person2)
	andrew.name = "Andrew"
	andrew.age = 35

	fmt.Println(formatString2(andrew))

	var t interface{}
	t = andrew

	if v, ok := t.(Person2); ok {
		fmt.Println(v, ok)
	}
}
