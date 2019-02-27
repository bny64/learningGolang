package main

import (
	"container/list"
	"fmt"
)

func main(){
	l := list.New()
	l.PushBack(10)
	l.PushBack(20)
	l.PushBack(30)

	fmt.Println("Front ", l.Front().Value)
	fmt.Println("Back " , l.Back().Value)

	for e := l.Front(); e != nil; e = e.Next(){
		fmt.Printf("%d ", e.Value)
	}

	for e := l.Back(); e != nil; e = e.Prev() {
		fmt.Printf("%d ", e.Value)
	}
}