package main

import (
	"fmt"
	"runtime"
	"sync"
)

func hello() {
	fmt.Println("Hello, world!")
}

func main3() {
	runtime.GOMAXPROCS(runtime.NumCPU() / 2)
	once := new(sync.Once)

	for i := 0; i < 3; i++ {
		go func(n int) {
			fmt.Println("goroutine : ", n)
			once.Do(hello)
		}(i)
	}

	fmt.Scanln()
}
