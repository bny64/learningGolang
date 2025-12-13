package main

import "fmt"

type Juice struct {
	name     string
	price    int
	category string
	taste    Taste
	state    State
}

func NewJuice(
	name string,
	price int,
	category string,
	taste Taste,
	state State,
) *Juice {
	return &Juice{
		name:     name,
		price:    price,
		category: category,
		taste:    taste,
		state:    state,
	}
}

func (c *Juice) Make() error {
	switch c.state {
	case MakeDone:
		return fmt.Errorf("%s 주스는 이미 만들어져 있습니다.", c.name)
	case Done:
		return fmt.Errorf("%s 주스는 이미 만들어져 고객에게 서빙되었습니다.", c.name)
	}
	c.state = MakeDone
	return nil
}

func (c *Juice) Package() error {
	switch c.state {
	case Waiting:
		return fmt.Errorf("%s 주스는 아직 준비되지 않았습니다.", c.name)
	case PackageDone:
		return fmt.Errorf("%s 주스는 이미 포장이 완료되었습니다.", c.name)
	case Done:
		return fmt.Errorf("%s 주스는 이미 고객에게 서빙되었습니다.", c.name)
	}
	c.state = PackageDone
	return nil
}

func (c *Juice) Pick() error {
	if c.state == Waiting {
		return fmt.Errorf("%s 주스는 아직 준비되지 않았습니다.", c.name)
	} else if c.state == MakeDone {
		return fmt.Errorf("%s 주스는 아직 포장되지 않았습니다.", c.name)
	} else if c.state == Done {
		return fmt.Errorf("%s 주스는 이미 고객에게 서빙되었습니다.", c.name)
	}
	c.state = Done
	return nil
}

func (c *Juice) Name() string {
	return c.name
}

func (c *Juice) Price() int {
	return c.price
}

func (c *Juice) Category() string {
	return c.category
}

func (c *Juice) Taste() Taste {
	return c.taste
}

func (c *Juice) State() State {
	return c.state
}
