package main

import "fmt"

type MyInt int

type Rectangle struct {
	width, height int
}

func (i MyInt) Print() {
	fmt.Println(i)
}

func (r Rectangle) Print() {
	fmt.Println(r.width, r.height)
}

type Printer interface {
	Print()
}

func main1() {
	var i MyInt = 5
	r := Rectangle{20, 30}

	var p Printer
	p = i
	p.Print()

	p = r
	p.Print()

	pArr := []Printer{i, r}
	for index := range pArr {
		pArr[index].Print()
	}

	for _, value := range pArr {
		value.Print()
	}

	r.Print()
	i.Print()

}
