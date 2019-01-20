package main

import "fmt"

type Person1 struct {
	name string
	age  int
}

type Student1 struct {
	Person1
	school string
	grade  int
}

func (p *Student1) greeting() {
	fmt.Println("Hello Student~")
}

func (p *Person1) greeting() {
	fmt.Println("Hello")
}

func funcOne(p Person1) {
	fmt.Println(p)
}

func funcTwo(p *Person1) {
	fmt.Println(p)
}

func (p Person1) funcThree(number int) {
	fmt.Println(p)  //{}
	fmt.Println(&p) //1456132
	fmt.Println(number)
}

func (p *Person1) funcFour(number int) {
	fmt.Println(p)  //&{}
	fmt.Println(&p) //{}
	fmt.Println(number)
}

func main() {
	var s Student1
	s.Person1.greeting()
	s.greeting()

	var pp Person1
	funcOne(pp)
	//funcTwo(pp)
	fmt.Println("-----------------")
	fmt.Println(&pp)
	fmt.Println("-----------------")
	pp.funcThree(10)
	pp.funcFour(1)
}
