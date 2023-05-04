package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	nums := []int64{2, 4, 6, 8, 10}
	var resultNums int64

	var wg sync.WaitGroup

	for _, value := range nums {
		wg.Add(1)
		go func(num int64) {
			//thread-safe operation to deal with a number
			atomic.AddInt64(&resultNums, num*num)
			wg.Done()
		}(value)
	}

	wg.Wait()

	fmt.Printf("Start data: %v\n Sum of squares equal to %v\n", nums, resultNums)
}
