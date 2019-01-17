package main

import "fmt"

func f2() {
	func() {
		s := recover()
		fmt.Println(s)
	}()
	a := [...]int{1, 2}

	for i := 0; i < 5; i++ {
		fmt.Println(a[i])
	}

}

func main2() {
	f2()
	fmt.Println("hello.world!")
}
