package main

import (
	"fmt"
	"unsafe"
)

func main1() {
	var r1 rune = '한'
	var r2 rune = '\U0000d55c'
	fmt.Println(r1, r2)

	var num1 int8 = 3 //the next of int is bit
	var num2 float32 = 2.2
	fmt.Println(float32(num1) + num2)

	//var num3 uint8 = math.MaxUint8 + 1
	//fmt.Println(num3)

	//the maximum of variable
	//standard is byte
	fmt.Println(unsafe.Sizeof(num2))
	fmt.Println(unsafe.Sizeof(num1))
}
