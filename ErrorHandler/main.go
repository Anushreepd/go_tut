package main

import "fmt"

func division(a, b float64) (float64, error) {
	if b <= 0 {
		return 0, fmt.Errorf("b cannot be zero")
	}
	return a / b, nil
}
func main() {
	fmt.Printf("Error handling")
	ans, _ := division(10, 2)
	fmt.Println("answer is ", ans)

}
