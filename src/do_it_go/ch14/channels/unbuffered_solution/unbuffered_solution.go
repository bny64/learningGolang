package main

import "fmt"

func sendValue(value int, channel chan int) {

	channel <- value
}

func main() {
	unbufferedChannel := make(chan int)

	go sendValue(1, unbufferedChannel)

	fmt.Println(<-unbufferedChannel)
}
