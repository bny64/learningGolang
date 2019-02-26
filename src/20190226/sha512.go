package main

import (
	"crypto/sha512"
	"fmt"
	"reflect"
)

func main(){
	s := "Hello, world!"
	h1 := sha512.Sum512([]byte(s))
	fmt.Printf("%x\n", h1) //16진수

	sha := sha512.New()
	sha.Write([]byte("Hello, "))
	sha.Write([]byte("world!"))
	h2 := sha.Sum(nil)
	fmt.Printf("%x\n", h2)
	fmt.Println(reflect.TypeOf(h1), reflect.TypeOf(h2))
	h3 := h1[:] //배열을 슬라이스로 변환.
	fmt.Println(reflect.TypeOf(h3)==reflect.TypeOf(h2))

}