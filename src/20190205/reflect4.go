package main

import (
	"fmt"
	"reflect"
)

func main4() {
	var a *int = new(int)
	*a = 1

	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.ValueOf(a))
	//fmt.Println(reflect.ValueOf(a).Int())
	fmt.Println(reflect.ValueOf(a).Elem())
	fmt.Println(reflect.ValueOf(a).Elem().Int())

	var b interface{}
	b = 1

	fmt.Println(reflect.TypeOf(b))        // int
	fmt.Println(reflect.ValueOf(b))       // <int Value>
	fmt.Println(reflect.ValueOf(b).Int()) // 1
}
