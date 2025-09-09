package problems

func CountOddNumbers(arr [6]int) int {
	count := 0
	for i := 0; i < len(arr); i++ {
		if arr[i]%2 != 0 {
			count++
		}
	}
	return count
}
