package main

import "sync/atomic"

// Счетчик с использованием атомарных операций
type AtomicCounter struct {
	counter int64
}

func NewAtomicCounter() *AtomicCounter {
	return &AtomicCounter{counter: 0}
}

func (c *AtomicCounter) Add() {
	atomic.AddInt64(&c.counter, 1)
}

func (c *AtomicCounter) Read() int64 {
	return atomic.LoadInt64(&c.counter)
}
