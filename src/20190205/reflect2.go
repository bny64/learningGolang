package main

import (
	"fmt"
	"reflect"
)

func main2() {
	var f float64 = 1.3
	t := reflect.TypeOf(f)
	v := reflect.ValueOf(f)

	fmt.Println(v)
	fmt.Println(t.Name())
	fmt.Println(t.Size())
	fmt.Println(t.Kind() == reflect.Float64)

	fmt.Println(t.Kind() == reflect.Int64)

	fmt.Println(v.Type())
	fmt.Println(v.Kind() == reflect.Float64)

	fmt.Println(v.Kind() == reflect.Int64)

	fmt.Println(v.Float())
}
