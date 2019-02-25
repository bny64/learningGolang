package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main(){
	s := "Hello, world!"
	err := ioutil.WriteFile("review.txt", []byte(s), os.FileMode(644))

	if err != nil {
		fmt.Println(err)
		return
	}

	data, err := ioutil.ReadFile("review.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(data))
}