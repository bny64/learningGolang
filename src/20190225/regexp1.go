package main

import (
	"fmt"
	"regexp"
)

func main(){
	re1, _ := regexp.Compile("\\.[a-zA-Z]+$")
	s1 := re1.ReplaceAllString("hello example.com", ".net")
	fmt.Println(s1) //정규표현식에 해당하는 문자열을 지정된 문자열로 변경

	re2, _ := regexp.Compile("([a-zA-Z,]+) ([a-zA-Z!]+)")
	s2 := re2.ReplaceAllString("Hello, world!", "${2}")
	fmt.Println(s2) //정규표현식에 해당하는 문자열을 지정된 문자열로 바꿈.

	re3, _ := regexp.Compile("([a-zA-Z,]+) ([a-zA-Z!]+)")
	s3 := re3.ReplaceAllString("Hello, world!", "${2} ${1}")
	fmt.Println(s3) //두 문자열의 위치를 바꿈.

	re4, _ := regexp.Compile("(?P<first>[a-zA-Z,]+) (?P<second>[a-zA-Z!]+)")
	s4 := re4.ReplaceAllString("Hello, world!", "${second} ${first}")
	fmt.Println(s4) //찾은 문자열에 ${forst} ${second}로 이름을 정함

	re5, _ := regexp.Compile("(?P<first>[a-zA-Z,]+) (?P<second>[a-zA-Z!]+)")
	s5 := re5.ReplaceAllLiteralString("Hello, world!", "${second} ${first}")
	fmt.Println(s5)
}