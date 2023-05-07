package main

// счетчик с использованием канала
type ChannelCounter struct {
	ch      chan func()
	counter int
}

// Горутина счетчика считывает все операции из канала и вызывает их поочереди
func NewChannelCounter() *ChannelCounter {
	counter := &ChannelCounter{counter: 0, ch: make(chan func(), 100)}
	go func(counter *ChannelCounter) {
		for f := range counter.ch {
			f()
		}
	}(counter)
	return counter
}

// Каждая операция вызываемая на данном счетчике ставится в очереди канала
// Все операции представлены в канале, в виде функции func()
func (c *ChannelCounter) Add() {
	c.ch <- func() {
		c.counter++
	}
}

func (c *ChannelCounter) Read() int {
	read := make(chan int)
	c.ch <- func() {
		read <- c.counter
		close(read)
	}
	return <-read
}
