package main

import (
	"fmt"
	"unicode/utf8"
)

func main2(){
	var s1 string = "한"
	fmt.Println(len(s1))

	var r1 rune = '한'
	fmt.Println(utf8.RuneLen(r1))

	var s2 string = "안녕하세요"
	fmt.Println(utf8.RuneCountInString(s2))

	b := []byte("안녕하세요")

	r, size := utf8.DecodeRune(b)
	fmt.Printf("%c, %d\n", r, size)

	r, size = utf8.DecodeRune(b[3:])
	fmt.Printf("%c, %d\n", r, size)

	r, size = utf8.DecodeLastRune(b[:len(b)-3])
	fmt.Printf("%c, %d\n", r, size)

	s3 := "Hello, world!"

	fmt.Printf("%c\n", s3[0])
	fmt.Printf("%c\n", s3[len(s3)-1])	

	r, _ =  utf8.DecodeRuneInString(s2)
	fmt.Printf("%c\n", r)

	r, _ = utf8.DecodeLastRuneInString(s2)
	fmt.Printf("%c\n", r)

	fmt.Println("---------------------------")

	b1 := []byte("안녕하세요")
	fmt.Println(utf8.Valid(b1))//utf8 확인
	b2 := []byte{0xff, 0xf1, 0xc1}
	fmt.Println(utf8.Valid(b2))

	var r2 rune = '한'
	fmt.Println(utf8.ValidRune(r2))
	var r3 rune = 0x11111111
	fmt.Println(utf8.ValidRune(r3))
	
	s4 := "한글"
	fmt.Println(utf8.ValidString(s4))
	s5 := string([]byte{0xff, 0xf1, 0xc1})
	fmt.Println(utf8.ValidString(s5))
}