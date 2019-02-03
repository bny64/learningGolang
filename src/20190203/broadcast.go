package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main2() {
	runtime.GOMAXPROCS(runtime.NumCPU() / 2)
	var mutex = new(sync.Mutex)
	var cond = sync.NewCond(mutex)

	c := make(chan bool, 3)

	for i := 0; i < 3; i++ {
		go func(n int) {
			mutex.Lock() //뮤텍스 잠금.
			c <- true    //채널 c에 true를 보냄
			fmt.Println("wait begin : ", n)
			cond.Wait() //조건변수 대기
			fmt.Println("wait end : ", n)
			mutex.Unlock() //뮤텍스 잠금 해제
		}(i)
	}

	for i := 0; i < 3; i++ {
		<-c //채널에서 값을 꺼냄
	}

	mutex.Lock() //뮤텍스 잠금
	fmt.Println("broadcast")
	cond.Broadcast()
	mutex.Unlock()

	fmt.Scanln()
}
