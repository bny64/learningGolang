package main

import "fmt"

type Taste string

const (
	Sweet       Taste = "SWEET"
	Bitter      Taste = "BITTER"
	FruitFlavor Taste = "FRUIT"
	Heavy       Taste = "HEAVY"
)

type Coffee struct {
	Name     string
	Price    int
	Category string
	Taste    Taste
}

func main() {
	coffee := new(Coffee)
	fmt.Print(coffee)
}
