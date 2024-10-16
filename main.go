package main

import (
	"fmt"
	"os"
	"sync"
)

func main() {
	commandLineArgs := os.Args[1:]

	if len(commandLineArgs) == 0 {
		fmt.Println("no website provided")
		os.Exit(1)
	} else if len(commandLineArgs) > 1 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	rawBaseURL := commandLineArgs[0]
	fmt.Printf("starting crawl of: %s\n", rawBaseURL)

	cfg := config{
		pages:              map[string]int{},
		rawBaseURL:         rawBaseURL,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, 5),
		wg:                 &sync.WaitGroup{},
	}

	cfg.crawlPage(rawBaseURL)

	fmt.Println(cfg.pages)
}
