package main

import "fmt"

type Building struct {
	Status string
}

func NewBuilding() *Building {
	return new(Building)
}

func (b *Building) Open() {
	b.Status = "OPEN"
}

func main() {
	building := NewBuilding()
	fmt.Printf("building status: %s\n", building.Status)

	building.Open()
	fmt.Printf("building status: %s\n", building.Status)
}
