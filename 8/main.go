package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	var number int64
	var bitIndex, bitValue int

	fSet := flag.NewFlagSet("set-bit", flag.ExitOnError)

	fSet.Int64Var(&number, "n", 20, "number int64")
	fSet.IntVar(&bitIndex, "i", 1, "index of bit (1-64)")
	fSet.IntVar(&bitValue, "b", 1, "value of bit (1 or 0)")

	if err := fSet.Parse(os.Args[1:]); err != nil {
		log.Fatalln(err)
	}

	//формирование битовой маски
	var mask int64 = 1 << (bitIndex - 1)
	if bitValue == 1 {
		fmt.Printf("original-%064b\n", number)
		fmt.Printf("bit mask-%064b\n", mask)

		//or операция
		number |= mask
		fmt.Printf("result---%064b\n", number)
		return
	}

	fmt.Printf("original-%064b\n", number)
	fmt.Printf("bit mask-%064b\n", mask)

	//and-not операция
	number &^= mask
	fmt.Printf("result---%064b\n", number)
}
