package main

import (
	"fmt"
	"strings"
)

func ReverseOrderOfWords(str string) string {
	ss, arrStr := strings.Split(strings.Trim(str, " "), " "), []string{}
	for i := len(ss) - 1; i >= 0; i-- {
		arrStr = append(arrStr, ss[i])
	}
	return strings.Join(arrStr, " ")
}

func main() {
	str := "sun dog down test"
	fmt.Printf("String before: %v\nString after: %v\n", str, ReverseOrderOfWords(str))
}
