package main

import "fmt"

func main3() {
	var num1 byte = 10
	var num2 byte = 0x32
	var num3 byte = 'a'

	fmt.Println(num1, num2, num3)
	//바이트는 문자열 저장 할 수 없음. 큰 따옴표 사용하면 에러
}
