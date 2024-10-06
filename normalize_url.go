package main

import (
	"net/url"
	"strings"
)

func normalizeURL(inputUrl string) (string, error) {
	parsedUrl, err := url.Parse(inputUrl)
	if err != nil {
		return "", err
	}

	return parsedUrl.Host + strings.TrimRight(parsedUrl.Path, "/"), nil
}
