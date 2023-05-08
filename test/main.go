package main

import (
	"fmt"
	"log"
)

func update(p *int) {
	b := 2
	p = &b
	log.Println(*p)
}

func main() {
	var (
		a = 1
		p = &a
	)
	fmt.Println(*p)
	update(p)
	fmt.Println(*p)

}
