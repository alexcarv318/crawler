package main

import (
	"fmt"
	"net/url"
)

func (cfg *config) crawlPage(rawCurrentURL string) {
	cfg.concurrencyControl <- struct{}{}
	defer func() {
		<-cfg.concurrencyControl
		cfg.wg.Done()
	}()

	if cfg.areMaxPagesReached() {
		return
	}

	if ok := checkIfURLDomainsAreEqual(cfg.rawBaseURL, rawCurrentURL); !ok {
		fmt.Printf("Error: URL domains mismatch - %s and %s\n", cfg.rawBaseURL, rawCurrentURL)
		return
	}

	normalizedCurrentURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error: URL normalization failed: %s\n", err)
		return
	}

	isFirstVisit := cfg.addPageVisit(normalizedCurrentURL)
	if !isFirstVisit {
		return
	}

	currentPageHTML, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error: Failed to get HTML from URL %s: %s\n", rawCurrentURL, err)
		return
	}

	currentPageContainedURLs, err := getURLsFromHTML(currentPageHTML, rawCurrentURL)
	if err != nil {
		fmt.Printf("Error: Failed to get links embedded into HTML: %s\n", err)
		return
	}

	fmt.Printf("from %s successfully crawled %d pages\n", normalizedCurrentURL, len(currentPageContainedURLs))

	// recursively go through all links in the website
	for _, nextURL := range currentPageContainedURLs {
		cfg.wg.Add(1)
		go cfg.crawlPage(nextURL)
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
