package main

import (
	"reflect"
	"testing"
)

func TestSortPages(t *testing.T) {
	testConfig, err := configure("https://example.com", 5, 10)
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name     string
		input    map[string]int
		expected map[string]int
	}{
		{
			name: "order count descending",
			input: map[string]int{
				"url1": 5,
				"url2": 1,
				"url3": 3,
				"url4": 10,
				"url5": 7,
			},
			expected: map[string]int{
				"url2": 1,
				"url3": 3,
				"url1": 5,
				"url5": 7,
				"url4": 10,
			},
		},
		{
			name: "alphabetize",
			input: map[string]int{
				"d": 1,
				"a": 1,
				"e": 1,
				"b": 1,
				"c": 1,
			},
			expected: map[string]int{
				"d": 1,
				"a": 1,
				"e": 1,
				"b": 1,
				"c": 1,
			},
		},
		{
			name:     "empty map",
			input:    map[string]int{},
			expected: map[string]int{},
		},
		{
			name:     "nil map",
			input:    nil,
			expected: map[string]int{},
		},
		{
			name: "one key",
			input: map[string]int{
				"url1": 1,
			},
			expected: map[string]int{
				"url1": 1,
			},
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			testConfig.pages = tc.input
			actual := testConfig.sortPagesByEmbeddedLinks()
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}
