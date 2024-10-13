package main

import (
	"fmt"
	"os"
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

	pages := map[string]int{}
	pages, err := crawlPage(rawBaseURL, rawBaseURL, pages)
	if err != nil {
		fmt.Printf("Unable to crawl page: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(pages)
}
