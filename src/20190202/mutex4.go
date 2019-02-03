package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main4() {
	runtime.GOMAXPROCS(runtime.NumCPU() / 2)

	data := 0
	var rwMutex = new(sync.RWMutex)

	go func() {
		for i := 0; i < 3; i++ {
			rwMutex.Lock()
			data += 1
			fmt.Println("write : ", data)
			time.Sleep(10 * time.Millisecond)
			rwMutex.Unlock()
		}
	}()

	go func() {
		for i := 0; i < 3; i++ {
			rwMutex.RLock()
			data += 1
			fmt.Println("read 1 : ", data)
			time.Sleep(1 * time.Second)
			rwMutex.RUnlock()
		}
	}()

	go func() {
		for i := 0; i < 3; i++ {
			rwMutex.RLock()
			fmt.Println("read 2 : ", data)
			time.Sleep(2 * time.Second)
			rwMutex.RUnlock()
		}
	}()

	time.Sleep(6 * time.Second)
}
