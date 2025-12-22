package main

import (
	"fmt"
	"time"
)

func longRunningTask() {
	for {
		fmt.Println("Long-running task is still running...")
		time.Sleep(1 * time.Second)
	}
}

func startTask() {
	go longRunningTask()
}

func main() {
	startTask()

	time.Sleep(2 * time.Second)
	fmt.Println("Main function ends.")
}
