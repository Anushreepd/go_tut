package main

import (
	"fmt"
	"my_learn/Array/problems"
)

func main() {
	fmt.Println("Sum of Even number")

	arr := [6]int{4, 7, 2, 9, 6, 1}
	fmt.Println("sum is", problems.SumEven(arr))
	fmt.Println("--------------------------------")

	value, index := problems.MaxElement(arr)
	fmt.Printf("max is %d, index is %d\n", value, index)

	fmt.Println("--------------------------------")

	fmt.Println("count of odd numbers are", problems.CountOddNumbers(arr))
	fmt.Println("--------------------------------")

	fmt.Println("reverse of array:", problems.ArrayReverse(arr))
	fmt.Println("--------------------------------")

	fmt.Println("It's", problems.CheckArraySorted(arr))

}
