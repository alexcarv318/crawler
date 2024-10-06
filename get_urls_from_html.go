package main

import (
	"golang.org/x/net/html"
	"log"
	"net/url"
	"strings"
)

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	var urls []string

	doc, err := html.Parse(strings.NewReader(htmlBody))
	if err != nil {
		log.Fatal(err)
	}

	for c := doc.FirstChild; c != nil; c = c.NextSibling {
		parsedHrefValue := getHrefValueFromAnchorNode(c)
		if parsedHrefValue == "" {
			continue
		}

		URL, err := addHostToURLIfNotExists(parsedHrefValue, rawBaseURL)
		if err != nil {
			return nil, err
		}

		urls = append(urls, URL)
	}

	return urls, nil
}

func getHrefValueFromAnchorNode(n *html.Node) string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				return a.Val
			}
		}
	}
	return ""
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
