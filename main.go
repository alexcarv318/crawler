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
	} else {
		fmt.Printf("starting crawl of: %s\n", commandLineArgs[0])
	}

	html, err := getHTML(commandLineArgs[0])
	if err != nil {
		fmt.Printf("Unable to get html from provided url: %v\n", err)
	}

	fmt.Println(html)
}
