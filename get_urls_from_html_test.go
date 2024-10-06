package main

import (
	"reflect"
	"testing"
)

func TestGetURLsFromHTML(t *testing.T) {
	type testCase struct {
		inputURL  string
		inputBody string
		expected  []string
	}

	t.Run("parse absolute and relative URLs", func(t *testing.T) {
		tc := testCase{
			inputURL: "https://google.com",
			inputBody: `
			<html>
				<body>
					<a href="/path/one">
						<span>Google.com</span>
					</a>
					<a href="https://other.com/path/one">
						<span>Google.com</span>
					</a>
				</body>
			</html>
			`,
			expected: []string{"https://google.com/path/one", "https://other.com/path/one"},
		}

		parsedURLs, err := getURLsFromHTML(tc.inputBody, tc.inputURL)
		if err != nil {
			t.Errorf("Failed to parse absolute and relative URLs: got unexpected error: %v", err)
		}

		if reflect.DeepEqual(parsedURLs, tc.expected) == false {
			t.Errorf("Failed to parse absolute and relative URLs: expected %v, got %v", tc.expected, parsedURLs)
		}
	})

	t.Run("all links was parsed", func(t *testing.T) {
		tc := testCase{
			inputURL: "https://google.com",
			inputBody: `
			<html>
				<body>
					<a href="/path/one">
						<span>Google.com</span>
					</a>
					<a href="/path/two">
						<span>Google.com</span>
					</a>
					<a href="/path/three">
						<span>Google.com</span>
					</a>
					<a href="/path/four">
						<span>Google.com</span>
					</a>
					<a href="/path/five">
						<span>Google.com</span>
					</a>
				</body>
			</html>
			`,
			expected: []string{
				"https://google.com/path/one",
				"https://google.com/path/two",
				"https://google.com/path/three",
				"https://google.com/path/four",
				"https://google.com/path/five",
			},
		}

		parsedURLs, err := getURLsFromHTML(tc.inputBody, tc.inputURL)
		if err != nil {
			t.Errorf("Failed to parse absolute and relative URLs: got unexpected error: %v", err)
		}

		if len(parsedURLs) != len(tc.expected) {
			t.Errorf("Failed to parse absolute and relative URLs: expected length %v, got %v", len(tc.expected), len(parsedURLs))
		}
	})

	t.Run("empty array if empty html body", func(t *testing.T) {
		tc := testCase{
			inputURL:  "https://google.com",
			inputBody: "",
			expected:  []string{"doesnt matter here"},
		}

		parsedURLs, err := getURLsFromHTML(tc.inputBody, tc.inputURL)
		if err != nil {
			t.Errorf("Failed to parse absolute and relative URLs: got unexpected error: %v", err)
		}
		if len(parsedURLs) > 0 {
			t.Errorf("Expected empty array of strings, got %v", parsedURLs)
		}
	})

	t.Run("empty array if there are <a> tags without href attribute", func(t *testing.T) {
		tc := testCase{
			inputURL: "https://google.com",
			inputBody: `
			<html>
				<body>
					<a>
						<span>Google.com</span>
					</a>
					<a>
						<span>Google.com</span>
					</a>
					<a>
						<span>Google.com</span>
					</a>
					<a>
						<span>Google.com</span>
					</a>
					<a>
						<span>Google.com</span>
					</a>
				</body>
			</html>
			`,
			expected: []string{"doesnt matter here"},
		}

		parsedURLs, err := getURLsFromHTML(tc.inputBody, tc.inputURL)
		if err != nil {
			t.Errorf("Failed to parse absolute and relative URLs: got unexpected error: %v", err)
		}
		if len(parsedURLs) > 0 {
			t.Errorf("Expected empty array of strings, got %v", parsedURLs)
		}
	})
}
