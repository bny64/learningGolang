package main

import "fmt"

type Person struct {
	Name string
}

func (p Person) Walk() {
	fmt.Println("걷고 있습니다.")
}

func (p Person) Talk() {
	fmt.Printf("안녕하세요 제 이름은 %s 입니다.\n", p.Name)
}

type Student struct {
	Person
}

func (s Student) Study() {
	fmt.Println("공부하고 있습니다.")
}

type Engineer struct {
	Person
}

func (e Engineer) Develop() {
	fmt.Println("개발하고 있습니다.")
}


