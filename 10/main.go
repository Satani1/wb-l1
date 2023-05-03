package main

import "fmt"

func main() {
	//стартовые данные температур
	temps := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	//создание мапы для группировки температурных значений
	tempGroups := make(map[int][]float64)

	for _, value := range temps {
		//вычисление температурного шага(ключа) для каждой температуры температуры, с последующим добавлением в слайс температур ключа
		tempTen := int(value/10) * 10
		tempGroups[tempTen] = append(tempGroups[tempTen], value)
	}

	fmt.Printf("Temperture data: %v\nResult temperture data: %v\n", temps, tempGroups)
}
