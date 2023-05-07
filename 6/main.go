package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Goroutines struct {
	wg *sync.WaitGroup
}

// закрытие канала
// Горутина читает значения из канала, если канал закрыт, то она завершит свою работу
func (g *Goroutines) closeChannel(ch <-chan int) {
	defer g.wg.Done()
	for {
		v, ok := <-ch
		if !ok {
			fmt.Printf("Finish close channel method\n\n")
			return
		}
		fmt.Printf("Got value - %v\n", v)
	}
}

// Завершение горутины с использованием сигнала остановки из контекста
func (g *Goroutines) closeCancelContext(ctx context.Context, ch chan int) {
	defer g.wg.Done()
	for {
		select {
		case v := <-ch:
			fmt.Printf("Got value - %v\n", v)
		case <-ctx.Done():
			fmt.Printf("Finish context cancel method\n\n")
			return
		}

	}
}

// Завершение с использованием сигнала остановки из контекста по времени
// если операция слишком долгая по времени, то горутина завершится
func (g *Goroutines) closeTimeoutContext(ctx context.Context, ch chan int) {
	defer g.wg.Done()
	for {
		select {
		//долгая операция
		case <-time.After(5 * time.Second):
			fmt.Printf("Got value - %v\n", <-ch)
		case <-ctx.Done():
			fmt.Printf("Finish context timeout method\n\n")
			return
		}
	}
}

// Использование отдельного канала для завершения
func (g *Goroutines) closeWithChannel(ch <-chan int, stop <-chan struct{}) {
	defer g.wg.Done()
	for {
		select {
		case v := <-ch:
			fmt.Printf("Got value - %v\n", v)
		case <-stop:
			fmt.Printf("Finish quit channel method\n\n")
			return
		}
	}
}
func main() {
	g := Goroutines{}

	//context timeout
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	chan1 := make(chan int, 1)

	g.wg = &sync.WaitGroup{}
	g.wg.Add(1)

	chan1 <- 1

	go g.closeTimeoutContext(ctx, chan1)

	close(chan1)
	g.wg.Wait()

	//context cancel
	ctx, cancel := context.WithCancel(context.Background())
	chan1 = make(chan int, 1)

	g.wg = &sync.WaitGroup{}
	g.wg.Add(1)

	chan1 <- 2

	go g.closeCancelContext(ctx, chan1)

	time.Sleep(2 * time.Second)
	cancel()
	close(chan1)

	g.wg.Wait()

	//closing channel
	chan1 = make(chan int, 2)
	chan1 <- 3
	chan1 <- 4

	g.wg = &sync.WaitGroup{}
	g.wg.Add(1)

	go g.closeChannel(chan1)

	time.Sleep(2 * time.Second)

	close(chan1)
	g.wg.Wait()

	//using stop channel
	chan1 = make(chan int, 1)
	stop := make(chan struct{})

	chan1 <- 5

	g.wg = &sync.WaitGroup{}
	g.wg.Add(1)

	go g.closeWithChannel(chan1, stop)

	time.Sleep(2 * time.Second)
	stop <- struct{}{}
	close(chan1)
	g.wg.Wait()

}
