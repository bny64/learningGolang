package main

import "fmt"

type Duck struct{}

type Person struct{}

func (d Duck) quack() {
	fmt.Println("quack")
}

func (p Person) quack() {
	fmt.Println("nope quack")
}

type Quacker interface {
	quack()
}

func inTheForest(q Quacker) {
	q.quack()
}

func main2() {
	var donald Duck
	var john Person

	quackers := []Quacker{donald, john}
	var quackers2 = make([]Quacker, 5)
	quackers2[0] = Quacker(donald)
	quackers2[1] = Quacker(john)

	for index := range quackers {
		quackers[index].quack()
		quackers2[index].quack()
	}
}
