package main

import "log"

func main() {
	var mySlice []string

	mySlice = append(mySlice, "John")
	mySlice = append(mySlice, "Trevor")
	mySlice = append(mySlice, "Samson")
	log.Println(mySlice)

	var isTrue bool

	isTrue = true

	if isTrue {
		log.Println("isTrue is", isTrue)
	} else {
		log.Println("isTrue is", isTrue)
	}

	myVar := "cat"

	switch myVar {
	case "cat":
		log.Println("My var is cat")
	case "dog":
		log.Println("My var is dog")
	case "bird":
		log.Println("My var is bird")
	default:
		log.Println("My var is something else")
	}
}
