package main

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/bny64/learningGolang/helpers"
)

const numPool = 10

func CalculateValue(intChan chan int) {
	randomNumber := helpers.RandomNumber(numPool)
	intChan <- randomNumber
}

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog    bool   `json:"has_dog"`
}

func main() {
	myJson := `
	[
		{
			"first_name": "John",
			"last_name": "Doe",
			"hair_color": "Brown",
			"has_dog": true
		},
		{
			"first_name": "Jane",
			"last_name": "Doe",
			"hair_color": "Black",
			"has_dog": false
		}
	]`

	var unmarshalled []Person

	err := json.Unmarshal([]byte(myJson), &unmarshalled)
	if err != nil {
		log.Println("Error unmarshalling", err)
	}

	log.Printf("unmarshalled: %v\n", unmarshalled)

	//wite json from a struct
	var mySlice []Person

	var m1 Person

	m1.FirstName = "Wally"
	m1.LastName = "Smith"
	m1.HairColor = "Black"
	m1.HasDog = true

	mySlice = append(mySlice, m1)

	m2 := Person{
		FirstName: "Bill",
		LastName:  "Williams",
		HairColor: "Brown",
		HasDog:    false,
	}

	mySlice = append(mySlice, m2)

	newJson, err := json.MarshalIndent(mySlice, "", "\t")
	if err != nil {
		log.Println("Error marshalling", err)
	}

	log.Println("newJson:", string(newJson))

	result, err := divide(100.0, 10.0)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("result:", result)
}

func divide(x, y float32) (float32, error) {
	if y == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return x / y, nil
}
