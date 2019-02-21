package main

import "fmt"

func main1(){
	s1 := "안녕하세요"
	r1 := []rune(s1)
	fmt.Println(string(r1[1]))

	for i, rune := range "Hello, 世界" {
		fmt.Printf("%d: %c\n", i, rune)
	}
}