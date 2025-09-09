package problems

func SumEven(arr [6]int) int {
	var sum = 0
	for i := 0; i < len(arr); i++ {
		if arr[i]%2 == 0 {
			sum = sum + arr[i]
		}
	}
	return sum
}
