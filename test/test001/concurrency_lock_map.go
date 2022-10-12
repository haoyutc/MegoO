package main

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type Ban struct {
	visitIps map[string]time.Time
	lock     sync.Mutex
}

func NewBan(ctx context.Context) *Ban {
	ban := &Ban{visitIps: make(map[string]time.Time)}
	go func() {
		timer := time.NewTimer(time.Minute * 1)
		for {
			select {
			case <-timer.C:
				ban.lock.Lock()
				for k, v := range ban.visitIps {
					if time.Now().Sub(v) >= time.Minute*1 {
						delete(ban.visitIps, k)
					}
				}
				ban.lock.Unlock()
				timer.Reset(time.Minute * 1)
			case <-ctx.Done():
				return
			}
		}
	}()
	return ban
}

func (ban *Ban) visit(ip string) bool {
	ban.lock.Lock()
	defer ban.lock.Unlock()
	if _, ok := ban.visitIps[ip]; ok {
		return true
	}
	ban.visitIps[ip] = time.Now()
	return false
}

func main() {
	success := int64(0)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ban := NewBan(ctx)
	wg := &sync.WaitGroup{}
	wg.Add(1000 * 100)
	for i := 0; i < 1000; i++ {
		for j := 0; j < 100; j++ {
			go func(j int) {
				defer wg.Done()
				ip := fmt.Sprintf("192.168.1.%d", j)
				if !ban.visit(ip) {
					atomic.AddInt64(&success, 1)
				}
			}(j)
		}
	}
	wg.Wait()
	fmt.Println("success:", success)
}
