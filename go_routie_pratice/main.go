package main

import (
	"fmt"
	"time"
)

// Write a function that prints numbers from 1 to 5.
// Run that function in a goroutine.
// What happens if the main function exits before the goroutine finishes?----> yes

/* func prints(i int) {
	fmt.Println(i)
} */

// Launch three goroutines, each printing a different message.
// Ensure the main function waits long enough for them to finish using time.Sleep().

/* func sayHi() {
	fmt.Println("hi func")
}

func sayHello() {
	fmt.Println("hello func")
} */

// Anonymous Functions
// Use an anonymous function inside a goroutine to print a message.

func main() {
	fmt.Println("routie problems")

	/* for i := 1; i <= 5; i++ {
		go prints(i)
		go sayHi()
		go sayHello()
	}
	time.Sleep(3000 * time.Millisecond)
	fmt.Println("done") */

	// Anonoymous func with parameters
	name := "hello go"
	go func(name string) {
		fmt.Println("this is anonymous func ", name)
	}(name)
	time.Sleep(3000 * time.Millisecond)
}
