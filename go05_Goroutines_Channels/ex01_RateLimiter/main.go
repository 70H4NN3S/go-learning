package main

import (
	"fmt"
	"sync"
	"time"
)

type RateLimiter struct {
	mu       sync.Mutex
	capacity int
	refill   float64
}

func NewRateLimiter(capacity int, refillPerSec float64) *RateLimiter {
	return &RateLimiter{
		sync.Mutex{},
		capacity,
		refillPerSec,
	}
}

func (r *RateLimiter) Allow() bool {
	return r.capacity > 0
}

func (r *RateLimiter) Wait() {
	r.mu.Lock()
	defer r.mu.Unlock()
}

func (r *RateLimiter) Increment() {
	for {
		time.Sleep(1 * time.Second)
		r.mu.Lock()
		r.capacity += int(r.refill)
		r.mu.Unlock()
	}
}

func (r *RateLimiter) Request() {
	if !r.Allow() {
		fmt.Println("waiting...")
		r.Wait()
	}
	fmt.Println("requesting....")
	r.capacity--
}

func main() {
	r := NewRateLimiter(5, 10)
	go r.Increment()
	for range 20 {
		r.Request()
	}
}
