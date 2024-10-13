package main

import (
	"fmt"
	"net/url"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) (map[string]int, error) {
	if ok := checkIfURLDomainsAreEqual(rawBaseURL, rawCurrentURL); !ok {
		return nil, fmt.Errorf("URL domains mismatch")
	}

	normalizedCurrentURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		return pages, err
	}

	// skip if visited
	if _, ok := pages[normalizedCurrentURL]; ok {
		pages[normalizedCurrentURL]++
		return pages, nil
	}

	pages[normalizedCurrentURL] = 1

	currentPageHTML, err := getHTML(rawCurrentURL)
	if err != nil {
		return pages, err
	}

	fmt.Println(currentPageHTML)

	currentPageContainedURLs, err := getURLsFromHTML(currentPageHTML, rawCurrentURL)
	if err != nil {
		return pages, err
	}

	// recursively go through all links in the website
	for _, nextURL := range currentPageContainedURLs {
		_, _ = crawlPage(rawBaseURL, nextURL, pages)
	}

	return pages, nil
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
