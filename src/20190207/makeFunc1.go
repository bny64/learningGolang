package main

import (
	"fmt"
	"reflect"
)

func h(args []reflect.Value) []reflect.Value {
	fmt.Println("Hello, world")
	return nil
}

func main1() {
	var hello func()                     //함수 담을 변수 선언
	fn := reflect.ValueOf(&hello).Elem() //hello 주소를 넘긴 뒤 Elem으로 값 정보를 가져온다.

	v := reflect.MakeFunc(fn.Type(), h)
	fn.Set(v) //hello의 값 정보인 fn에 h의 함수 정보 v를 설정하여 함수를 연결.

	hello()
}
