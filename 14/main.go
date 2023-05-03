package main

import (
	"fmt"
	"reflect"
)

func main() {
	ch := make(chan any)
	dataTypes := []any{16, "test string", true, ch}

	for _, elem := range dataTypes {
		v := reflect.ValueOf(elem)
		fmt.Printf("'%v' has %s type\n", elem, v.Kind().String())
	}
}
