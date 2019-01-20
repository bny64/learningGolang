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

func main() {
	var s Student1
	s.Person1.greeting()
	s.greeting()
}
