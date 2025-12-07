package main

import "fmt"

func main() {
	var value1 uint8 = 187 //10111011
	var value2 uint8 = 222 //11011110

	// var shiftValue uint8 = 19
	// var shiftLength uint8 = 2

	fmt.Println("AND result")
	fmt.Println("========")
	andResult := value1 & value2
	fmt.Printf("%08b & %08b = %08b\n", value1, value2, andResult)
}
