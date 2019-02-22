package main

import (
	ft "fmt"
	sv "strconv"
	rf "reflect"
)

func main4(){
	var s1 string
	s1 = sv.FormatBool(true)
	ft.Println(s1, rf.TypeOf(s1))

	var s2 string
	s2 = sv.FormatFloat(1.3, 'f', -1, 32)//변환 실수 값, 실수 형식, 실수의 정밀도, 부동소수점 비트수
	ft.Println(s2)

	var s3 string
	s3 = sv.FormatInt(-10, 10)
	ft.Println(s3)

	var s4 string
	s4 = sv.FormatUint(32, 16)
	ft.Println(s4)

	var i1 = 100
	var i2 int64 = 100
	ft.Println(sv.Itoa(i1))
	ft.Println(sv.FormatInt(i2, 10))
}