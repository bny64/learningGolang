package main

import (
	f "fmt"
	s "strings"
	u "unicode"
)

func main2(){
	f.Println(s.Trim("Hello, ~~ world! ~~", "~")) //일치하는 것 모두 제거
	f.Println(s.TrimSuffix("Hello, ~~ world! ~~", "~")) //해당 문자열로 끝나는 것만 제거
	//s.TrimLeft
	//s.TrimRight

	fc := func(r rune) bool {
		return u.Is(u.Hangul, r)
	}

	var s1 = "안녕 Hello 고 Go 언어"
	f.Println(s.TrimFunc(s1, fc))
	f.Println(s.TrimLeftFunc(s1, fc))
	f.Println(s.TrimRightFunc(s1, fc))
	
	f.Println(s.TrimRight(s1, "언"))

	r := s.NewReplacer("<", "&lt;", ">", "&gt;")
	f.Println(r.Replace("<div><span>Hello, world!</span></div>"))
}