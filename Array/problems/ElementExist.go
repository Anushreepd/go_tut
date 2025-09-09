package problems

func ElementExists(arr [6]int, ele int) bool {
	for i := 0; i < len(arr); i++ {
		if ele == arr[i] {
			return true
		}
	}
	return false
}
