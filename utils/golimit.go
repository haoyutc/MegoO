package utils

// 限流器

type Limiter struct {
	ch chan int
}

func NewLimiter(max int) *Limiter {
	return &Limiter{ch: make(chan int, max)}
}

func (g *Limiter) Add() {
	g.ch <- 1
}

func (g *Limiter) Done() {
	<-g.ch
}

func (g *Limiter) Close() {
	close(g.ch)
}
