package middleware

import (
	"sync"
	"testing"
)

func TestRateLimiterConcurrentAllow(t *testing.T) {
	rl := NewRateLimiter(100000)
	start := make(chan struct{})
	var wg sync.WaitGroup
	const workers = 4
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			<-start
			for j := 0; j < 300; j++ {
				rl.Allow("10.0.0.1")
			}
		}()
	}
	close(start)
	wg.Wait()
	if !rl.Allow("10.0.0.1") {
		t.Fatal("unexpected token exhaustion")
	}
}

func TestRateLimiterConcurrentStatsReset(t *testing.T) {
	rl := NewRateLimiter(1000)
	start := make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		<-start
		for j := 0; j < 500; j++ {
			rl.Stats()
		}
	}()
	go func() {
		defer wg.Done()
		<-start
		for j := 0; j < 500; j++ {
			rl.Reset("10.0.0.2")
		}
	}()
	close(start)
	wg.Wait()
}

func TestRateLimiterConcurrentPeek(t *testing.T) {
	rl := NewRateLimiter(1000)
	start := make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		<-start
		for j := 0; j < 500; j++ {
			rl.Allow("10.0.0.3")
		}
	}()
	go func() {
		defer wg.Done()
		<-start
		for j := 0; j < 500; j++ {
			rl.Peek("10.0.0.3")
		}
	}()
	close(start)
	wg.Wait()
}

func TestRateLimiterSnapshotNoEscape(t *testing.T) {
	rl := NewRateLimiter(1000)
	rl.Allow("10.0.0.1")
	snap := rl.Snapshot()
	snap["10.0.0.1"].tokens = -999
	if rl.Peek("10.0.0.1") == -999 {
		t.Fatal("snapshot escaped internal reference")
	}
}
