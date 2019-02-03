package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main(){
	runtime.GOMAXPROCS(runtime.NumCPU()/2)

	 var mutex = new(sync.Mutex)
	 var cond = sync.NewCond(mutex)

	 c:=make(chan bool, 3) //buffer 3 channal

	 for i:=0; i<3; i++{
		go func(n int){
			mutex.Lock()
			c <- true
			fmt.Println("wait begin : ", n)
			cond.Wait()
			fmt.Println("wait end : ", n)
			mutex.Unlock()
		}(i)
	 }

	 for i:=0; i<3; i++ {
		 <- c//채널에서 값을 꺼냄. 고루틴 3개가 모두 실행될 때 까지 ㄱ디ㅏ림.
	 }

	 for i:=0; i<3; i++ {
		 mutex.Lock()
		 fmt.Println("signal : " , i)
		 cond.Signal()
		 mutex.Unlock()

	 }

	 fmt.Scanln()
}