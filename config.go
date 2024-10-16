package main

import (
	"sync"
)

type config struct {
	pages              map[string]int
	rawBaseURL         string
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
}
