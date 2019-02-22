package main

import (
	f "fmt"
	s "strings"
	u "unicode"
)

func main1(){
	s1 := []string{"Hello, ", "world!"}
	f.Println(s.Join(s1, " "))

	s2 := s.Split("Hello, world!", " ")
	f.Println(s2[1])

	s3 := s.Fields("Hello, world!")
	f.Println(s3[1])

	ff := func(r rune) bool {
		return u.Is(u.Hangul, r)
	}

	s4 := s.FieldsFunc("Hello안녕Hello", ff)
	f.Println(s4)

	f.Println(s.Repeat("hello",3))

	f.Println(s.Replace("Hello, world!", "world", "Go", 1))
}