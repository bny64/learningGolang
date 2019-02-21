package main

import (
	f "fmt"
	s "strings"
	"unicode"
)

func main(){
	str1 := "Hello, world!"
	str2 := "반갑습니다. 반가운"
	f.Println(s.Contains(str1, "wo"))
	f.Println(s.ContainsAny(str1, "lhrld"))
	f.Println(s.Count(str1, "o"))
	f.Println("--------")
	f.Println(s.ContainsRune(str2, '반'))
	f.Println(s.HasPrefix(str1, "He"))
	f.Println(s.HasSuffix(str2, "운"))
	f.Println(s.HasSuffix(str2, "반"))
	f.Println("--------")
	f.Println(s.Index(str1, "rl"))
	f.Println(s.Index(str1, "hw"))
	f.Println(s.IndexAny(str1, "lwd"))
	f.Println("--------")
	f.Println(s.IndexByte(str1, 'l'))
	f.Println(s.IndexRune(str2, '가'))


	ff := func(r rune) bool {
		return unicode.Is(unicode.Hangul, r) // r이 한글 유니코드이면 true를 리턴
	}

	f.Println(s.IndexFunc(str2, ff))
}