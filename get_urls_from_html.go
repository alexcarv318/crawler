package main

import (
	"golang.org/x/net/html"
	"net/url"
	"strings"
)

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	var urls []string

	doc, err := html.Parse(strings.NewReader(htmlBody))
	if err != nil {
		return urls, err
	}

	// recursive function to go deeper through html nodes and search for <a> tags
	// done according to official docs https://pkg.go.dev/golang.org/x/net/html#example-Parse
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					if correctURL, err := addHostToURLIfNotExists(a.Val, rawBaseURL); err == nil {
						urls = append(urls, correctURL)
						break
					}
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}

	f(doc)
	return urls, nil
}

func addHostToURLIfNotExists(inputURL, host string) (string, error) {
	parsedURL, err := url.Parse(inputURL)
	if err != nil {
		return "", err
	}

	if parsedURL.Host == "" {
		return host + inputURL, nil
	}

	return inputURL, nil
}
