package main

type Job struct {
	ID      int
	Payload interface{}
}

type Result struct {
	JobID  int
	Output interface{}
	Err    error
}

type Pool struct{}

func NewPool(workers int) *Pool {
	return nil
}

func (p *Pool) Submit(job Job) {
}

func Results() <-chan Result {
	return nil
}

func Stop() {
}
