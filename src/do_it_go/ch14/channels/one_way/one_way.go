package main

import "fmt"

func sendChannel(value int, channel chan<- int) {
	channel <- value
}

func receiveChannel(channel <-chan int) {
	fmt.Printf("채널 수신: %d\n	", <-channel)
}

func main() {
	channel := make(chan int, 1)

	//값 송신
	sendChannel(1, channel)

	//값 수신
	receiveChannel(channel)
}
