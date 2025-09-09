package problems

import "fmt"

func MaxElement(arr [6]int) (int, int) {
	fmt.Println("maximum value")
	max := arr[0]
	index := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
			index = i
		}
	}
	return max, index
}
