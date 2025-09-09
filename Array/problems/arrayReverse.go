package problems

func ArrayReverse(arr [6]int) [6]int {
	left := 0
	right := len(arr) - 1

	for left < right {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}
	return arr
}
