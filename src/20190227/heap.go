package main

import (
	"fmt"
	"container/heap"
)

type MinHeap []int

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	r := h[i] < h[j]
	return r
}

func (h MinHeap) Swap(i, j int){
	fmt.Printf("Swap %d %d\n", h[i], h[j])
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x interface{}){
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{}{
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n -1]
	return x
}

func main(){
	data := new(MinHeap)

	heap.Init(data)
	heap.Push(data, 5)
	heap.Push(data, 2)
	heap.Push(data, 7)
	heap.Push(data, 3)

	fmt.Println(data, "최솟값 : ", (*data)[0])
	
}