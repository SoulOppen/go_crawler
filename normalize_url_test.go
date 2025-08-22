package main

import "testing"

func TestNormalizeURL(t *testing.T) {
	tests := []struct {
		name     string
		inputURL string
		expected string
	}{
		{
			name:     "remove https scheme",
			inputURL: "https://blog.boot.dev/path",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "remove http scheme",
			inputURL: "http://example.com/abc",
			expected: "example.com/abc",
		},

		{
			name:     "remove scheme no path",
			inputURL: "https://example.com",
			expected: "example.com",
		},
		{
			name:     "with www subdomain",
			inputURL: "http://www.test.com/home",
			expected: "www.test.com/home",
		},
		{
			name:     "with query string",
			inputURL: "https://example.com/search?q=golang",
			expected: "example.com/search?q=golang",
		},
		{
			name:     "with fragment",
			inputURL: "https://docs.site.com/page#section2",
			expected: "docs.site.com/page#section2",
		},
		{
			name:     "with port number",
			inputURL: "http://localhost:8080/api",
			expected: "localhost:8080/api",
		},
		{
			name:     "IP address",
			inputURL: "https://127.0.0.1/login",
			expected: "127.0.0.1/login",
		},

		{
			name:     "empty string",
			inputURL: "",
			expected: "",
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := normalizeURL(tc.inputURL)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			}
			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}
