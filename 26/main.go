package main

import (
	"fmt"
	"strings"
)

func IsUnique(str string) bool {
	str = strings.ToLower(str)
	mChars := make(map[rune]struct{})
	for _, char := range str {
		if _, ok := mChars[char]; ok {
			return false
		}
		mChars[char] = struct{}{}
	}
	return true
}

func main() {
	arrStr := []string{"abcd", "abCdefAaf", "aabcd", "aaaa", "empty"}

	for _, str := range arrStr {
		fmt.Printf("String: %v - %v\n", str, IsUnique(str))
	}
}
