package main

import "fmt"

func sum(a, b int) (res int) {
	return a + b
}

func main() {
	fmt.Println("function use")
	fmt.Println("sum is --->", sum(4, 5))
}
