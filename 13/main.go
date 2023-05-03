package main

import "fmt"

func main() {
	a, b := 1, 2
	fmt.Printf("a - %v\nb - %v\n\n", a, b)

	a, b = b, a
	fmt.Printf("a - %v\nb - %v\n\n", a, b)
}
