package main

import (
	ft "fmt"
	sv "strconv"
	rf "reflect"
)

func main3(){
	var err error

	var num1 int
	num1, err = sv.Atoi("100")
	ft.Println(num1, err)

	var num2 int
	num2, err = sv.Atoi("10t")
	ft.Println(num2, err)

	var s string
	s = sv.Itoa(50)
	ft.Println(s, rf.TypeOf(s))

}