package main

import (
	"fmt"
	"sort"
)

func main(){
	a := []int{10, 5, 3, 7, 6}
	b := []float64{4.2, 7.6, 5.5, 1.3, 9.9}
	c := []string{"Maria", "Andrew", "John"}
	d := make([]int, len(a), cap(a))
	_ = b
	_ = c

	sort.Sort(sort.IntSlice(a)) //int슬라이스 오름차 순
	fmt.Println(a)
	sort.Sort(sort.Reverse(sort.IntSlice(a))) //int 슬라이스 내림차 순
	fmt.Println(a)

	fmt.Printf("cap : %d, len : %d\n", cap(d), len(d))
	//copy(d, a[:])
	copy(d, a)
	sort.Ints(d)
	fmt.Println(d)

	sort.Sort(sort.Float64Slice(b)) //float64 슬라이스 오름차 순
	fmt.Println(b)
	sort.Sort(sort.Reverse(sort.Float64Slice(b))) //float64 슬라이스 내림차 순
	fmt.Println(b)

	sort.Sort(sort.StringSlice(c)) //string 슬라이스 오름차 순
	fmt.Println(c)
	sort.Sort(sort.Reverse(sort.StringSlice(c))) //string 슬라이스 내림차 순
	fmt.Println(c)
}