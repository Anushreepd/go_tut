package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	IsAdult bool   `json:"is_adult"`
}

func main() {
	person1 := Person{Name: "anu", Age: 24, IsAdult: true}
	//marshal (encoding)
	json_marshal, err := json.Marshal(person1)
	if err != nil {
		fmt.Println("error while unmarshaling data")
	}
	fmt.Println("json marshal data", string(json_marshal))

	//umarshal
	var person2 Person
	error := json.Unmarshal(json_marshal, &person2)
	if error != nil {
		fmt.Println("error while unmarshaling data")
	}
	fmt.Println("json Unmarshal data", person2)
}
