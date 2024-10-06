package main

import "testing"

func TestNormalizeURL(t *testing.T) {
	tests := []struct {
		name     string
		inputUrl string
		expected string
	}{
		{
			name:     "remove secure scheme",
			inputUrl: "https://www.google.com/search",
			expected: "www.google.com/search",
		},
		{
			name:     "remove scheme",
			inputUrl: "http://www.google.com/search",
			expected: "www.google.com/search",
		},
		{
			name:     "remove secure scheme and slash in the end",
			inputUrl: "https://www.google.com/search/",
			expected: "www.google.com/search",
		},
		{
			name:     "remove scheme and slash in the end",
			inputUrl: "http://www.google.com/search/",
			expected: "www.google.com/search",
		},
		{
			name:     "remove slash in the end",
			inputUrl: "www.google.com/search/",
			expected: "www.google.com/search",
		},
		{
			name:     "don't do nothing",
			inputUrl: "www.google.com/search",
			expected: "www.google.com/search",
		},
		{
			name:     "return the same if not an URL",
			inputUrl: "not-an-url",
			expected: "not-an-url",
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := normalizeURL(tc.inputUrl)
			if err != nil {
				t.Errorf("Test failed %v (%s): got unexpected error: %v.", i, tc.name, err)
			}

			if actual != tc.expected {
				t.Errorf("Test failed %v (%s): expected %v, actual %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}
