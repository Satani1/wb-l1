package main

import (
	"fmt"
)

func BinarySearch(arr []int, item int) int {
	low := 0
	high := len(arr)
	for low < high {
		mid := (low + high) / 2
		guess := arr[mid]
		if guess == item {
			return mid
		} else if guess > item {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

func main() {
	integers := []int{-20, -15, 2, 5, 10, 31, 45, 100}
	fmt.Printf("Index of search item: %v\n", BinarySearch(integers, 2))
}
