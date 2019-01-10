package main

import (
	"fmt"
	"unicode/utf8"
)

func main2() {
	var s1 = "한글"
	s2 := "Hello"

	fmt.Println(s1, s2)
	fmt.Println(len(s1)) //6

	fmt.Println(utf8.RuneCountInString(s1)) //when we count a real length of utf-8

	fmt.Println(s1[0])
	fmt.Printf("%c\n", s1[1]) // we cannot revise string in a golang
}
