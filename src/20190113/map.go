package main

import "fmt"

//before5
func main6() {
	var firstMap map[string]int
	//할당
	//make(map[키_자료형]값_자료형)
	var secondMap = make(map[string]int)
	thirdMap := make(map[string]int)
	sixthMap := map[string]int{}
	_, _, _, _ = firstMap, secondMap, thirdMap, sixthMap

	//선언과 할당을 동시에
	forthMap := map[string]int{"Hello": 10, "world": 20}
	fifthMap := map[string]int{
		"Hello": 10,
		"world": 20,
	}

	fmt.Println(forthMap["Hello"])
	fmt.Println(fifthMap["world"])
}
