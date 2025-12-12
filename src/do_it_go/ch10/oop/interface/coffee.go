package main

import "fmt"

type Coffee struct {
	name     string
	price    int
	category string
	taste    Taste
	state    State
}

func NewCoffee(
	name string,
	price int,
	category string,
	taste Taste,
	state State) *Coffee {
	coffee := new(Coffee)
	coffee.name = name
	coffee.price = price
	coffee.category = category
	coffee.taste = taste
	coffee.state = state
	return coffee
}

func (c *Coffee) Make() error {
	if c.state == MakeDone {
		return fmt.Errorf("%s 커피는 이미 만들어져 있습니다.", c.name)
	} else if c.state == Done {
		return fmt.Errorf("%s 커피는 이미 만들어져 고객에게 서빙되었습니다.", c.name)
	}
	c.state = MakeDone
	return nil
}

func (c *Coffee) Package() error {
	if c.state == Waiting {
		return fmt.Errorf("%s 커피는 아직 준비되지 않았습니다.", c.name)
	} else if c.state == PackageDone {
		return fmt.Errorf("%s 커피는 이미 포장이 완료되었습니다.", c.name)
	} else if c.state == Done {
		return fmt.Errorf("%s 커피는 이미 고객에게 서빙되었습니다.", c.name)
	}
	c.state = PackageDone
	return nil
}
