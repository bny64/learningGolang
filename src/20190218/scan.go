package main

import "fmt"

func main() {
	var num1, num2 int
	n, _ := fmt.Scanf("%d, %d", &num1, &num2)
	fmt.Printf("입력 갯수 : %d\n", n)
	fmt.Printf("%d, %d\n", num1, num2)
}
