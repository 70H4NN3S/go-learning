package main

import (
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
		time.Sleep(time.Second)
		r.mu.Lock()
		r.capacity += int(r.refill)
		r.mu.Unlock()
	}
}

func (r *RateLimiter) Request() {
	if !r.Allow() {
		r.Wait()
	}
	r.capacity--
}

func main() {
}
