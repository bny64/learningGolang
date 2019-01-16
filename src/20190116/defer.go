package main

import "fmt"

func hello2() {
	fmt.Println("Hello")
}

func world() {
	fmt.Println("world")
}

//before 4
func main() {
	//defer은 현재 함수가 끝나기 바로 직전에 호출됨
	defer world()
	hello2()
	hello2()
	hello2()

	HelloWorld()
}

func HelloWorld() {
	defer func() {
		fmt.Println("@world")
	}()

	func() {
		fmt.Println("@Hello")
	}()

	deferFunc()
}

func deferFunc() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d", i)
	}
}
