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

func NewCoffee(
	name string,
	price int,
	category string,
	taste Taste) *Coffee {
	coffee := new(Coffee)
	coffee.Name = name
	coffee.Price = price
	coffee.Category = category
	coffee.Taste = taste
	return coffee

}

func main() {
	americano := NewCoffee("아메리카노", 3000, "블랜딩 커피", Bitter)
	latte := NewCoffee("카페라떼", 3500, "블랜딩 카페", Sweet)

	fmt.Printf("americano: %v\n", americano)
	fmt.Printf("latte: %v\n", latte)

}
