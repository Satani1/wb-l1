package main

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"time"
)

// Воркер читает число из канала и выводит его в stdout.
func RunWorker(ctx context.Context, wg *sync.WaitGroup, input <-chan int, nWorker int) {
	defer wg.Done()
	for {
		select {
		case n := <-input:
			fmt.Printf("%v worker read --- %v\n", nWorker, n)
		case <-ctx.Done():
			fmt.Printf("%v worker killed, because %v\n", nWorker, ctx.Err())
			return
		}
	}
}

func main() {
	numsWorkers := flag.Int("workers", 10, "Number of workers")
	flag.Parse()

	var wg sync.WaitGroup

	inputChan := make(chan int, 5)
	defer close(inputChan)
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	//запуск n-го количества воркеров
	for i := 0; i < *numsWorkers; i++ {
		//start worker
		wg.Add(1)
		go RunWorker(ctx, &wg, inputChan, i)
	}

	//постоянная запись данных в канал.
	//Используется канал ticker для интервалов времени между записью
	go func() {
		wg.Add(1)
		defer wg.Done()

		for {
			select {
			case <-ticker.C:
				a := rand.Intn(100)
				inputChan <- a
				fmt.Printf("Writer send to channel: %v\n", a)
			case <-ctx.Done():
				fmt.Printf("Exit from writer\n")
				return
			}
		}
	}()

	//после нажатия ctrl+c программа начнет завершатся.
	//Сигнал отправится на все горутины через контекст,
	//после чего ждем их завершения.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	cancel()
	wg.Wait()
	fmt.Println("Exit from programm")

}
