package main

import "fmt"

func main() {
	data := []string{"cat", "cat", "dog", "cat", "tree"}

	mapCheck := make(map[string]struct{})
	resultData := make([]string, 0)

	//Инетируемся по изначальному слайсу данных и проверяем, есть ли он в мапе(как ключ).
	//Если есть, то пропускаем это значение и интерируемся далее(пропуская строки 17-18, иначе мы добавляем это значение в слайс и в мапу
	for _, value := range data {
		if _, ok := mapCheck[value]; ok {
			continue
		}
		resultData = append(resultData, value)
		mapCheck[value] = struct{}{}
	}

	fmt.Printf("Starting data: %v\n Result set: %v\n", data, resultData)
}
