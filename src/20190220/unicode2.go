package main

import (
	f "fmt"
	"unicode"
)

func main2(){
	f.Println(unicode.IsGraphic('1'))
	f.Println(unicode.IsGraphic('a'))
	f.Println(unicode.IsGraphic('한'))
	f.Println(unicode.IsGraphic('韓'))
	
	f.Println(unicode.IsLetter('a'))
	f.Println(unicode.IsLetter('1'))

	f.Println(unicode.IsDigit('1'))
	f.Println(unicode.IsControl('\n'))
	f.Println(unicode.IsMark('\u17c9'))
	
	f.Println(unicode.IsPrint('1'))
	f.Println(unicode.IsPunct('.'))

	f.Println(unicode.IsSpace(' '))   // true: ' '는 공백이므로 true
	f.Println(unicode.IsSymbol('♥')) // true: ♥는 심볼이므로 true

	f.Println(unicode.IsUpper('A')) // true: A는 대문자이므로 true
	f.Println(unicode.IsLower('a')) // true: a는 소문자이므로 true
}