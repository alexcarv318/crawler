package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	commandLineArgs := os.Args[1:]

	if len(commandLineArgs) < 3 {
		fmt.Println("not enough arguments provided")
		fmt.Println("usage: crawler <baseURL> <maxConcurrency> <maxPages>")
		return
	}
	if len(commandLineArgs) > 3 {
		fmt.Println("too many arguments provided")
		return
	}

	rawBaseURL := commandLineArgs[0]
	maxConcurrencyString := commandLineArgs[1]
	maxPagesString := commandLineArgs[2]

	maxConcurrency, err := strconv.Atoi(maxConcurrencyString)
	if err != nil {
		fmt.Println("as max number of goroutines, please provide a valid integer")
	}

	maxPages, err := strconv.Atoi(maxPagesString)
	if err != nil {
		fmt.Println("as max number of pages to crawl, please provide a valid integer")
	}
	cfg, err := configure(rawBaseURL, maxConcurrency, maxPages)
	if err != nil {
		fmt.Printf("error configuring crawler: %s\n", err)
	}

	cfg.wg.Add(1)
	go cfg.crawlPage(rawBaseURL)
	cfg.wg.Wait()

	cfg.printReport(rawBaseURL)
}
