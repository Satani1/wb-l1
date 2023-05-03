package main

import "fmt"

func ReverseString(inputStr string) string {
	runes := []rune(inputStr)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func main() {
	stroke := "Its a test"

	fmt.Printf("Default string: %v\nReversed string: %v\n", stroke, ReverseString(stroke))
}
