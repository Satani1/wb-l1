package main

import (
	"fmt"
	"sync"
)

func main() {
	//Mutex Counter
	wg := &sync.WaitGroup{}
	mutexCounter := NewMCounter()

	numberIteration := 10

	wg.Add(numberIteration)
	for i := 0; i < numberIteration; i++ {
		go func() {
			defer wg.Done()
			mutexCounter.Add()
		}()
	}

	wg.Wait()
	fmt.Printf("Mutext counter - %v\n", mutexCounter.Read())

	//Atomic Counter
	wg = &sync.WaitGroup{}
	atomicCounter := NewAtomicCounter()

	wg.Add(numberIteration * 2)
	for i := 0; i < numberIteration*2; i++ {
		go func() {
			defer wg.Done()
			atomicCounter.Add()
		}()
	}

	wg.Wait()
	fmt.Printf("Atomic counter - %v\n", atomicCounter.Read())

	//Channel Counter
	wg = &sync.WaitGroup{}
	channelCounter := NewChannelCounter()

	wg.Add(numberIteration * 3)
	for i := 0; i < numberIteration*3; i++ {
		go func() {
			defer wg.Done()
			channelCounter.Add()
		}()
	}

	wg.Wait()
	fmt.Printf("Channels counter - %v\n", channelCounter.Read())
}
