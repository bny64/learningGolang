package main

import "fmt"

func main1() {
	r := sum(1, 2)
	fmt.Println(r)

	sum, diff := SumAndDiff(4, 5)
	fmt.Println(sum, diff)

	a, _, c, _, e := hello()
	fmt.Println(a, c, e)

	sum2, diff2 := SumAndDiff2(4, 5)
	fmt.Println(sum2, diff2)
}

func sum(a int, b int) int {
	return a + b
}

func sum2(a int, b int) (r int) {
	//리턴 값 이름 지정 가능
	//리턴 값 이름 지정시 return 뒤에 값 넣어주지 않아도 됨.
	r = a + b
	return
}

// SumAndDiff what
func SumAndDiff(a int, b int) (int, int) {
	return a + b, a - b
}

func hello() (int, int, int, int, int) {
	return 1, 2, 3, 4, 5
}

func SumAndDiff2(a int, b int) (sum int, diff int) {
	sum = a + b
	diff = a - b
	return
}
