package main

import (
	"fmt"
	"sync"
	"time"
)

func longRunningTask(stop chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-stop:
			//종료 신호를 받으면 반복을 종료
			fmt.Println("Long-running task is stopping...")
			return
		default:
			fmt.Println("Long-running task is still running...")
			time.Sleep(1 * time.Second)
		}
	}
}

func startTask(stop chan struct{}, wg *sync.WaitGroup) {
	wg.Add(1)
	go longRunningTask(stop, wg)
}

func main() {
	var wg sync.WaitGroup
	stop := make(chan struct{})

	startTask(stop, &wg)

	time.Sleep(2 * time.Second)
	fmt.Println("Main function ends.")

	close(stop)

	wg.Wait()
	fmt.Println("All tasks completed")
}
