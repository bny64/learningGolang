package main

import (
	"log"
	"time"
)

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

func main() {
	user := User{
		FirstName:   "Trevor",
		LastName:    "Smith",
		PhoneNumber: "1234567890",
		Age:         30,
		BirthDate:   time.Now(),
	}

	log.Println(user.FirstName)
	log.Println(user.LastName)
	log.Println(user.PhoneNumber)
	log.Println(user.Age)
	log.Println(user.BirthDate)

	user2 := &User{
		FirstName:   "Trevor",
		LastName:    "Smith",
		PhoneNumber: "1234567890",
		Age:         30,
		BirthDate:   time.Now(),
	}
	log.Println(user2.FirstName)

	myMap := make(map[string]string)

	myMap["dog"] = "Samson"
	myMap["other-dog"] = "Cassie"
	log.Println(myMap["dog"])
	log.Println(myMap["other-dog"])
	log.Println(myMap)
}

func saySomething(s string) (string, string) {
	return s, "World"
}
