package main

import (
	"fmt"
	"reflect"
)

type Data struct {
	a, b int
}

func main1() {
	num := 1
	fmt.Println(reflect.TypeOf(num))

	s := "Hello, world!"
	fmt.Println(reflect.TypeOf(s))

	f := 1.3
	fmt.Println(reflect.TypeOf(f))

	var data = Data{1, 2}
	fmt.Println(reflect.TypeOf(data))

}
