package main

import "fmt"

func main() {
	//안녕
	fmt.Println("안녕하세요.")

	var total = 1

	total = total + 3

	fmt.Println(total)

	message := "hello dear readers"
	fmt.Println(message)

	var (
		name    = "gil"
		age     = 30
		address = "seoul"
	)

	var msg string = "power"

	fmt.Println(name, age, address)
	fmt.Println(msg)
}
