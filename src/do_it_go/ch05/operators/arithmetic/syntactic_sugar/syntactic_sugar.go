package main

import "fmt"

func main() {
	operand := 5
	sum := 10

	sum = sum + operand //15
	fmt.Println(sum)

	sumSyntacticSugar := 10
	sumSyntacticSugar += operand
	fmt.Println(sumSyntacticSugar)

	difference := 10
	difference -= 5
	fmt.Println(difference)
}
