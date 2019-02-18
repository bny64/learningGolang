package main

import "fmt"

func main2() {
	var num1 = 10
	var num2 = 3.2
	var num3 complex64 = 2.5 + 8.1i
	var s string = "Hello, world!"
	var b bool = true
	var a []int = []int{1, 2, 3}
	var m map[string]int = map[string]int{"Hello": 1}
	var p *int = new(int)
	type Data struct{ a, b int }
	var data Data = Data{1, 2}
	var i interface{} = 1

	fmt.Println(num1)
	fmt.Println(num2)
	fmt.Println(num3)
	fmt.Println(s)
	fmt.Println(b)
	fmt.Println(a)
	fmt.Println(m)
	fmt.Println(p)
	fmt.Println(data)
	fmt.Println(i)
}
