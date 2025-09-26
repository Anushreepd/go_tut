package main

import (
	"fmt"
	"time"
)

func sayHi() {
	fmt.Println("hi function")
	time.Sleep(1000 * time.Millisecond)
	fmt.Print("hi function ended")
}
func sayHello() {
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("hello function")
}
func main() {
	fmt.Println("welcome to go routie")
	go sayHi()
	go sayHello()
	time.Sleep(6000 * time.Millisecond)
}
