package main

import "testing"

func TestNormalizeURL(t *testing.T) {
	tests := []struct {
		name     string
		inputURL string
		expected string
	}{
		{
			name:     "remove secure scheme",
			inputURL: "https://www.google.com/search",
			expected: "www.google.com/search",
		},
		{
			name:     "remove scheme",
			inputURL: "http://www.google.com/search",
			expected: "www.google.com/search",
		},
		{
			name:     "remove secure scheme and slash in the end",
			inputURL: "https://www.google.com/search/",
			expected: "www.google.com/search",
		},
		{
			name:     "remove scheme and slash in the end",
			inputURL: "http://www.google.com/search/",
			expected: "www.google.com/search",
		},
		{
			name:     "remove slash in the end",
			inputURL: "www.google.com/search/",
			expected: "www.google.com/search",
		},
		{
			name:     "don't do nothing",
			inputURL: "www.google.com/search",
			expected: "www.google.com/search",
		},
		{
			name:     "return the same if not an URL",
			inputURL: "not-an-url",
			expected: "not-an-url",
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := normalizeURL(tc.inputURL)
			if err != nil {
				t.Errorf("Test failed %v (%s): got unexpected error: %v.", i, tc.name, err)
			}

			if actual != tc.expected {
				t.Errorf("Test failed %v (%s): expected %v, actual %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}
