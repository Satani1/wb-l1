package main

import "fmt"

func SetIntersection(firstArr, secondArr []int) []int {
	//создание мапы
	mapCheck := make(map[int]int)

	//каждое значение из 1-го и 2-го слайса добавляем в мапу, как ключ, и вместе с этим инкрементируем количество одинаковых значений.
	for _, value := range firstArr {
		mapCheck[value] += 1
	}
	for _, value := range secondArr {
		mapCheck[value] += 1
	}

	result := make([]int, 0)
	//Если значения повторяются в двух слайсах, то мы добавляем в новый слайс это значение и возвращаем его.
	for key, value := range mapCheck {
		if value > 1 {
			result = append(result, key)
		}
	}

	return result
}

func main() {
	firstSet := []int{1, 4, 3, 3, 5, -1}
	secondSet := []int{3, 1, -1, 20}

	resultSet := SetIntersection(firstSet, secondSet)

	fmt.Printf("First set: %v\nSecond set: %v\n Set intersection: %v\n", firstSet, secondSet, resultSet)
}
