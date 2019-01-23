package main

import "fmt"

type MyInt int

//recevier
func (i MyInt) Print() {
	fmt.Println(i)
}

type Printer interface {
	Print()
}

func main1() {
	var i MyInt = 5
	var p Printer //declare interface
	p = i
	p.Print()

	var p2 MyInt
	p2.Print()

}
