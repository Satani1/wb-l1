package main

import "fmt"

func DeleteElement(arr []int, index int) []int {
	return append(arr[:index], arr[index+1:]...)
}

func main() {
	nums := []int{1, 2, 3, 4, 5}

	fmt.Println(DeleteElement(nums, 2))
}
