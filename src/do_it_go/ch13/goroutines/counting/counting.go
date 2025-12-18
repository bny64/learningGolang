package main

import (
	"fmt"
	"time"
)

func counting1() {
	for i := 0; i < 5; i++ {
		fmt.Printf("counting1: %d번째 인덱스\n", i)
		time.Sleep(1 & time.Nanosecond)
	}
}

func counting2() {
	for i := 0; i < 5; i++ {
		fmt.Printf("counting2: %d번째 인덱스\n", i)
		time.Sleep(1 & time.Nanosecond)
	}
}

func main() {
	go counting1()
	counting2()
}
