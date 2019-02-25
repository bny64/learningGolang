package main

import (
	"fmt"
	"encoding/json"
	"reflect"
)

func main(){
	doc := `
	{
		"name":"maria",
		"age":10
	}
	`

	var data map[string]interface{} //json문서의 데이터를 저장할 공간을 맵으로 선언

	json.Unmarshal([]byte(doc), &data)
	fmt.Println(data["name"], data["age"])

	data = make(map[string]interface{}) //모든 자료형을 저장할 수 있는 맵 생성

	data["name"] = "maria"
	data["age"] = 10

	fmt.Println(reflect.TypeOf(doc))
	doc2, _ := json.Marshal(data)
	fmt.Println(string(doc2))
	fmt.Println(reflect.TypeOf(doc2))

	doc2, _ = json.MarshalIndent(data, "", "  ")
	fmt.Println(reflect.TypeOf(doc2))
	fmt.Println(string(doc2))
}