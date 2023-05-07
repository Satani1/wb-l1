package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

var justString string

func createHugeString(size int) string {
	//используем Builder для эффективной конкатенации строк
	var hugeStr strings.Builder
	for i := 0; i < size; i++ {
		hugeStr.WriteRune('г')
	}
	return hugeStr.String()
}

func someFunc() {
	v := createHugeString(1 << 10)

	fmt.Printf("Символ '%v' занимает %v байт\n", "г", utf8.RuneLen('г')) // руна занимает 2 байта
	//срез происходит по байтам
	justString = v[:100]
	fmt.Println(justString)

	//преобразуем нашу строку в слайс рун
	//выведем срез по рунам
	justRunes := []rune(v)
	justString = string(justRunes[:100])
	fmt.Println(justString)

}

func main() {
	someFunc()
}
