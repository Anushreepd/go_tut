package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type ToDo struct {
	UserId    int    `json:"userid"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func GetTodoList() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("error", err)
		return
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("error", err)
		return
	}
	fmt.Println(string(data))

	//unamrshal
	var todo ToDo
	error := json.Unmarshal(data, &todo)
	if error != nil {
		fmt.Println("error while unmarshaling", error)
	}
	fmt.Println(todo)
}

func PostTodoList() {
	todo := ToDo{
		UserId:    1,
		Title:     "read book",
		Completed: false,
	}

	json_marshal, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("error while marshaling data", err)
		return
	}
	jsonString := string(json_marshal)
	jsonReader := strings.NewReader(jsonString)
	myurl := "https://jsonplaceholder.typicode.com/todos"
	res, err := http.Post(myurl, "application/json", jsonReader)
	if err != nil {
		fmt.Println("error while grtting data", err)
		return
	}
	fmt.Println(res.Status)
}
func main() {
	fmt.Println("crud operation")
	//GetTodoList()
	//PostTodoList()
}
