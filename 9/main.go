package main

import "fmt"

// запись чисел из массива в канал
func ReadData(arr []int) <-chan int {
	resultChan := make(chan int)
	go func() {
		for _, value := range arr {
			resultChan <- value
		}
		close(resultChan)
	}()
	return resultChan
}

func SquareNums(inputChan <-chan int) <-chan int {
	resultChan := make(chan int)
	go func() {
		for value := range inputChan {
			resultChan <- value * value
		}
		close(resultChan)
	}()
	return resultChan
}

func main() {
	nums := []int{1, 2, 3, 4, 5}

	numsChannel := ReadData(nums)

	outputChannel := SquareNums(numsChannel)

	//вывод данных из канала
	for value := range outputChannel {
		fmt.Println(value)
	}

}
