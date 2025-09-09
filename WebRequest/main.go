package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("web request")
	res, err := http.Get("https://fakestoreapi.com/products/1")
	if err != nil {
		fmt.Println("error while fetching data")
	}
	defer res.Body.Close()
	//fmt.Println("respone", res) // http.response

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("error while getting data")
	}
	fmt.Println(string(data))
}
