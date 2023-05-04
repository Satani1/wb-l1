package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func RunWorker(wg *sync.WaitGroup, inputChan <-chan int) {
	defer wg.Done()
	for value := range inputChan {
		fmt.Printf("Read from channel: %v\n", value)
	}
	fmt.Printf("Worker has done.\n")
}

func main() {
	var wg sync.WaitGroup

	duration := flag.Duration("n", 10*time.Second, "Work duration of programm")
	flag.Parse()

	dataChan := make(chan int, 1)

	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()
	timer := time.NewTimer(*duration)

	//вызов воркера, который читает данные из канала и выводит их в stdout
	wg.Add(1)
	go RunWorker(&wg, dataChan)

	//Постоянная запись случайных чисел в канал dataChan с переодичностью
	//Как только завершится время duration канал закроется и программа завершится
	go func() {
		wg.Add(1)
		defer wg.Done()
		for {
			select {
			case <-ticker.C:
				a := rand.Intn(100)
				dataChan <- a
				fmt.Printf("Write to channel: %v\n", a)
			case <-timer.C:
				fmt.Printf("The programm closes, because time is up\n")
				close(dataChan)
				return
			}
		}
	}()

	wg.Wait()
}
