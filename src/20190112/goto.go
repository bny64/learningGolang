package main

import "fmt"

func main2() {
	var a = 1

	if a == 1 {
		goto ERROR1
	}

	if a == 2 {
		goto ERROR2
	}

	if a == 3 {
		goto ERROR1
	}

	fmt.Println(a)
	fmt.Println("Success")

ERROR1:
	fmt.Println("error1")
	return

ERROR2:
	fmt.Println("ERROR2")
	return

}
