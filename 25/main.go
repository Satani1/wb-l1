package main

import (
	"fmt"
	"time"
)

func Sleep(timer time.Duration) {
	<-time.After(timer)
}
func main() {
	fmt.Println("Start program")
	startTime := time.Now()

	Sleep(2 * time.Second)

	fmt.Printf("End program \nRun time = %v\n", time.Since(startTime))
}
