package main

import (
	"sync"
)

type Job struct {
	ID      int
	Payload any
}

type Result struct {
	JobID  int
	Output any
	Err    error
}

type Pool struct {
	Input     chan Job
	Output    chan Result
	wg        sync.WaitGroup
	processor func(Job) Result
}

func NewPool(workers int, processor func(Job) Result) *Pool {
	input := make(chan Job, 5)
	output := make(chan Result, 5)
	p := &Pool{input, output, sync.WaitGroup{}, processor}

	for range workers {
		p.wg.Add(1)
		go p.Worker()
	}
	return p
}

func (p *Pool) Worker() {
	defer p.wg.Done()
	for job := range p.Input {
		result := p.processor(job)
		p.Output <- result
	}
}

func (p *Pool) Submit(job Job) {
	p.Input <- job
}

func (p *Pool) Results() <-chan Result {
	return p.Output
}

func (p *Pool) Stop() {
	close(p.Input)
	p.wg.Wait()
	close(p.Output)
}
