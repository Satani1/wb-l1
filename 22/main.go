package main

import (
	"fmt"
	"log"
	"math/big"
)

func main() {
	num1, ok := new(big.Int).SetString("100000000000", 10)
	if !ok {
		log.Println("Cant create first big number")
	}

	num2, ok := new(big.Int).SetString("1000000000", 10)
	if !ok {
		log.Println("Cant create first big number")
	}

	result := big.NewInt(0)

	fmt.Printf("First number - %v\nSecond number - %v\n", num1, num2)

	fmt.Printf("Multiplacation = %v\n", result.Mul(num1, num2))
	fmt.Printf("Division = %v\n", result.Div(num1, num2))
	fmt.Printf("Summ = %v\n", result.Add(num1, num2))
	fmt.Printf("Substract = %v\n", result.Sub(num1, num2))
}
