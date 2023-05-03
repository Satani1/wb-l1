package main

import "fmt"

func QuickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	pivot := arr[0]
	less := make([]int, 0)
	greater := make([]int, 0)

	for _, element := range arr[1:] {
		if element <= pivot {
			less = append(less, element)
		} else {
			greater = append(greater, element)
		}
	}

	arr = append([]int{}, QuickSort(less)...)
	arr = append(arr, pivot)
	arr = append(arr, QuickSort(greater)...)
	return arr

}

func main() {
	Unsorted := []int{5, 1, -20, 10, 40, 55, 5}
	fmt.Printf("Unsorted array: %v\nAfter sorting:%v\n", Unsorted, QuickSort(Unsorted))
}
