package main

import "fmt"

type MyInt2 int

func (i MyInt2) Print2() {
	fmt.Println(i)
}

type Rectangle struct {
	width, height int
}

func (r Rectangle) Print2() {
	fmt.Println(r.width, r.height)
}

type Printer2 interface {
	Print2()
}

func main() {
	var i MyInt2 = 5 //MyInt type
	r := Rectangle{10, 20}
	var p Printer2 //인터페이스 선언
	p = i
	p.Print2()

	p = r
	p.Print2()
}
