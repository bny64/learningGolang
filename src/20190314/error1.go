package main

import (
	"fmt"
	"log"
)

func helloOne(n int) (string, error) {
	if n == 1 {
		s := fmt.Sprintln("Hello", n)
		return s, nil
	}

	return "", fmt.Errorf("%d는 1이 아닙니다.", n)
}

func main() {
	s, err := helloOne(1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(s)

	s, err = helloOne(2)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(s)

	fmt.Println("Hello, world!")
}
