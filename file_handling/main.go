package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// creare a file

	/* file_create, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("error while creating a file", err)
		return
	}
	fmt.Println("successsfully created a file")
	defer file_create.Close()

	//add data inside file
	content := "hello welcome to go programming"
	_, errors := io.WriteString(file_create, content)
	if errors != nil {
		fmt.Println("error while writing a file", err)
		return
	} */

	//read using buffer
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("error while creating a file", err)
		return
	}
	defer file.Close()

	buffer := make([]byte, 1024)

	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("error while reading a file", err)
			return
		}
		fmt.Println(string(buffer[:n]))
	}

}
