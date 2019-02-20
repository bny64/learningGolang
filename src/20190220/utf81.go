package main

import (
	f "fmt"
	"unicode/utf8"
)

func main3(){
	var s string = "한"
	var s2 string = "한글"
	f.Println(len(s))
	f.Println(len(s2))

	var r rune = '한'
	f.Println(utf8.RuneLen(r))

	f.Println(utf8.RuneCountInString(s2))
}