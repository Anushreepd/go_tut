package main

import (
	"fmt"
	"sync"
)

func worker(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("worker %d started\n", i)
	fmt.Printf("worker %d ended\n", i)
}

func main() {
	var wg sync.WaitGroup
	fmt.Println("sync worker")

	for i := 1; i <= 3; i++ {
		go worker(i, &wg)
		wg.Add(1)
	}

	wg.Wait()
	fmt.Println("All worker compeleted")
}
