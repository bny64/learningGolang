package main

import "fmt"

func f() {
	defer func() {
		s := recover()
		//panic함수에서 발생한 에러 메세지를 받아온다.
		//recover함수는 반드시 지연호출로 사용
		fmt.Println(s)
	}()

	panic("Error!")
}

func main1() {
	a := [...]int{1, 2}
	fmt.Println(a)

	//패닉이 발생하면 패닉이 발생된 곳에서 멈춤.
	fmt.Println("hello")

	f()
	fmt.Println("world!")
}
