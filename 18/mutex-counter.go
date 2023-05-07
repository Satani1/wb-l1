package main

import "sync"

// счетчик с использованием мьютекса
type MCounter struct {
	mu      *sync.RWMutex
	counter int
}

func NewMCounter() *MCounter {
	return &MCounter{counter: 0, mu: &sync.RWMutex{}}
}

// Во время записи одна горотуни блокирует остальные для записи
func (c *MCounter) Add() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counter++
}

// во время чтения блокируются горутины Add(), но горутины Read() также могут работать параллельно
func (c *MCounter) Read() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.counter
}
