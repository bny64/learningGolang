package main

import "fmt"

func main5() {
	for i := 1; i <= 100; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Println("3,5 공배수")
		case i%3 == 0:
			fmt.Println("3의 배수")
		case i%5 == 0:
			fmt.Println("5의 배수")
		default:
			fmt.Println(i)
		}
	}
}
