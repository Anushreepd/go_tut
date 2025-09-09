package problems

func CheckArraySorted(arr [6]int) string {
	for i := 0; i < len(arr); i++ {
		if arr[i] > arr[i+1] {
			return "Not a sorted Array"
		}
	}
	return " a sorted Array"
}
