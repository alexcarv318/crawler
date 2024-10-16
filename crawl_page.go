package main

import (
	"fmt"
	"net/url"
)

func (cfg *config) crawlPage(rawCurrentURL string) {
	if ok := checkIfURLDomainsAreEqual(cfg.rawBaseURL, rawCurrentURL); !ok {
		fmt.Printf("Error: URL domains mismatch - %s and %s\n", cfg.rawBaseURL, rawCurrentURL)
		return
	}

	normalizedCurrentURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error: URL normalization failed: %s\n", err)
		return
	}

	// skip if visited
	if _, ok := cfg.pages[normalizedCurrentURL]; ok {
		cfg.pages[normalizedCurrentURL]++
		return
	}

	cfg.pages[normalizedCurrentURL] = 1

	currentPageHTML, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error: Failed to get HTML from URL %s: %s\n", rawCurrentURL, err)
		return
	}

	fmt.Println(currentPageHTML)

	currentPageContainedURLs, err := getURLsFromHTML(currentPageHTML, rawCurrentURL)
	if err != nil {
		fmt.Printf("Error: Failed to get links embedded into HTML: %s\n", err)
		return
	}

	// recursively go through all links in the website
	for _, nextURL := range currentPageContainedURLs {
		cfg.crawlPage(nextURL)
	}
}

func checkIfURLDomainsAreEqual(url1, url2 string) bool {
	parsedBaseURL, err := url.Parse(url1)
	if err != nil {
		return false
	}
	parsedCurrentURL, err := url.Parse(url2)
	if err != nil {
		return false
	}

	if parsedBaseURL.Hostname() != parsedCurrentURL.Hostname() {
		return false
	} else {
		return true
	}
}
