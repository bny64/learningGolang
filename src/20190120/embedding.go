package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p *Person) greeting1() {

	fmt.Println(p)  //&{0}
	fmt.Println(&p) //0xc000086018
	fmt.Println("Hello~")
}

func (p Person) greeting2() {
	fmt.Println(p)  //{0}
	fmt.Println(&p) //&{0}
	fmt.Println("Hello~")
}

type Student struct {
	//p Persion -> has a 관계
	Person // -> is a 관계
	school string
	grade  int
}

func test1(p Student) {
	fmt.Println(p)
}

func main4() {
	var s Student
	//s.p.greeting()
	s.greeting1()
	s.greeting2()

	test1(s)
}
