package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	nums := []int{2, 4, 6, 8, 10}

	for _, value := range nums {
		wg.Add(1)

		go func(number int) {
			defer wg.Done()
			fmt.Printf("Square of '%v' equal to '%v'\n", number, number*number)
		}(value)
	}

	wg.Wait()

}
