package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	name string `tag1:"이름" tag2:"Name"`
	age  int    `tag1:"나이" tag2:"Age"`
}

type Earth struct {
	name string `plantName:"행성이름"`
}

func main3() {
	p := Person{}
	e := Earth{}

	name, ok := reflect.TypeOf(p).FieldByName("name")
	fmt.Println(ok, name.Tag.Get("tag1"), name.Tag.Get("tag2"))

	age, ok := reflect.TypeOf(p).FieldByName("age")
	fmt.Println(ok, age.Tag.Get("tag1"), age.Tag.Get("tag2"))

	plantName, ok := reflect.TypeOf(e).FieldByName("name")
	fmt.Println(plantName.Tag.Get("plantName"))
}
