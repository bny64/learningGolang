package main

import "fmt"

func hello(n *int) {
	*n = 2
}

//포인터 변수는 참조연산이기 때문에
//메모리 주소에 들어있는 값을 바꾸면 바뀜.
func main() {
	var n int = 1
	hello(&n)
	fmt.Println(n)
}
