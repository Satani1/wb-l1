package main

import "fmt"

// с сохранением порядка элементов
func DeleteElement(arr []int, index int) []int {
	//возвращаем слайс без элемента, который нам необходимо удалить
	return append(arr[:index], arr[index+1:]...)
}

func DeleteElement2(arr []int, index int) []int {
	newSlice := make([]int, 0)
	newSlice = append(newSlice, arr[:index]...)
	return append(newSlice, arr[index+1:]...)
}

// без сохранения порядка
func DeleteElementWithoutOrder(arr []int, index int) []int {
	//меняем элемент, который необходимо удалить, на первый элемент слайса
	//и возвращаем слайс без первого элемента
	arr[index] = arr[0]
	return arr[1:]
}

func DeleteElementWothoutOrder2(arr []int, index int) []int {
	//меняем элемент, который необходимо удалить, на последний элемент слайса
	//и возвращаем слайс без последнего элемента
	arr[index] = arr[len(arr)-1]
	return arr[:len(arr)-1]
}

// С использованием дженериков
func DeleteGeneric[T comparable](arr []T, index int) []T {
	return append(arr[:index], arr[index+1:]...)
}
func main() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Printf("Original slice: %v\n\n", nums)

	fmt.Printf("Deleting with order save: %v\n", DeleteElement(nums, 2))

	nums = []int{1, 2, 3, 4, 5}
	fmt.Printf("With preserving original slice (re-slice method): %v\n\n", DeleteElement2(nums, 2))

	nums = []int{1, 2, 3, 4, 5}
	fmt.Printf("Delete without original order. First method: %v\n", DeleteElementWithoutOrder(nums, 2))

	nums = []int{1, 2, 3, 4, 5}
	fmt.Printf("Second method: %v\n\n", DeleteElementWothoutOrder2(nums, 2))

	numsF := []float64{1.5, 2.3, 3.2, 4.9, 5.1}
	fmt.Printf("With generic method: %v\n", DeleteGeneric(numsF, 2))
}
