package main

import (
	"reflect"
	"testing"
)

func TestGetURLsFromHTML(t *testing.T) {
	tests := []struct {
		name      string
		inputURL  string
		inputBody string
		expected  []string
	}{
		{
			name:     "absolute and relative URLs",
			inputURL: "https://blog.boot.dev",
			inputBody: `
<html>
	<body>
		<a href="/path/one">
			<span>Boot.dev</span>
		</a>
		<a href="https://other.com/path/one">
			<span>Boot.dev</span>
		</a>
	</body>
</html>
`,
			expected: []string{"https://blog.boot.dev/path/one", "https://other.com/path/one"},
		},
		{
			name:     "relative URL with fragment",
			inputURL: "https://example.com",
			inputBody: `
<html>
	<body>
		<a href="/page#section1">Section</a>
	</body>
</html>
`,
			expected: []string{"https://example.com/page#section1"},
		},
		{
			name:     "relative URL with query",
			inputURL: "https://mysite.com",
			inputBody: `
<html>
	<body>
		<a href="/search?q=golang">Search</a>
		<a href="https://external.com/path?q=1">External</a>
	</body>
</html>
`,
			expected: []string{"https://mysite.com/search?q=golang", "https://external.com/path?q=1"},
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := getURLsFromHTML(tc.inputBody, tc.inputURL)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			}
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}
