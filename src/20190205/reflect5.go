package main

import (
	"fmt"
	"reflect"
)

func h(args []reflect.Value) []reflect.Value {
	fmt.Println("Hello, world!")
	return nil
}

func main(){
	var hello func()

	fn := reflect.ValueOf(&hello).Elem()

	v:=reflect.MakeFunc(fn.Type(), h) //h의 함수 정보 생성

	fn.Set(v)
	hello()
}