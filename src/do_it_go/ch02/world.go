package main

import "fmt"

func main() {
	var whatSay string = saySometing()
	fmt.Println("The function returned", whatSay)
	saySometing()
}

func saySometing() string {
	return "hello"
}
