package main

import "fmt"

//before6
func main7() {
	solarSystem := map[string]float32{}

	solarSystem["Mercury"] = 87.969
	solarSystem["Venus"] = 224.70069
	solarSystem["Earth"] = 365.25641

	if value, ok := solarSystem["Venus"]; ok {
		fmt.Println(value)
	}

	if value, ok := solarSystem["Uranus"]; ok {
		fmt.Println(value)
	}

	for key, value := range solarSystem {
		fmt.Println(key, value)
	}

	//delete map element
	//delete(맵, 삭제할 키)
	place := map[string]string{}
	place["Gangwondo"] = "wonju"
	place["chungnam"] = "choenan"
	place["kyungbuk"] = "busan"
	place["jeolla"] = "gwangju"

	fmt.Println(place)
	delete(place, "kyungbuk")
	fmt.Println(place)

}
