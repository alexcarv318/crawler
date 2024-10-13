package main

import (
	"errors"
	"io"
	"net/http"
)

func getHTML(rawURL string) (string, error) {
	resp, err := http.Get(rawURL)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", errors.New("error while getting HTML: status: " + resp.Status)
	}

	if resp.Header.Get("Content-Type") != "text/html" {
		return "", errors.New("error while getting HTML: the requested content-type is not HTML")
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
