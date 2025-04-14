package security_test

import (
	"bluebase/security"
	"fmt"
	"testing"
)

func TestSanitize(t *testing.T) {
	scenarios := []struct {
		content       string
		expectContent string
	}{
		{"<img src=\"http://url.to.file.which/not.exist\" onerror=alert(document.cookie);>", "<img src=\"http://url.to.file.which/not.exist\">"},
		{"<script>alert('XSS');</script>", ""},
		{"<a href=\"javascript:alert('XSS')\">Click me</a>", "Click me"},
		{"<div onclick=\"alert('XSS')\">Content</div>", "<div>Content</div>"},
		{"<p style=\"color:red;\">Styled Text</p>", "<p>Styled Text</p>"},
		{"<table><tr><td>Data</td></tr></table>", "<table><tr><td>Data</td></tr></table>"},
		{"<ul><li>Item 1</li><li>Item 2</li></ul>", "<ul><li>Item 1</li><li>Item 2</li></ul>"},
		{"<iframe src=\"http://example.com\"></iframe>", ""},
		{"<img src=\"javascript:alert('XSS')\">", ""}, {"<body onload=alert('XSS')>", ""},
		{"<b onmouseover=alert('XSS')>Bold</b>", "<b>Bold</b>"},
		{"<svg onload=alert('XSS')>", ""},
		{"<math><mi xlink:href=\"javascript:alert('XSS')\">X</mi></math>", "X"},
		{"<img src=x onerror=alert('XSS')>", "<img src=\"x\">"},
		{"<img src=\"data:image/svg+xml;base64,PHN2ZyBvbmxvYWQ9YWxlcnQoJ1hTUycpPg==\">", ""},
		{"<input type=\"text\" value=\"\" onfocus=\"alert('XSS')\">", ""},
		{"<form action=\"javascript:alert('XSS')\">", ""},
		{"<meta http-equiv=\"refresh\" content=\"0;url=javascript:alert('XSS');\">", ""},
		{"<link rel=\"stylesheet\" href=\"javascript:alert('XSS')\">", ""},
		{"<object data=\"javascript:alert('XSS')\"></object>", ""},
		{"<embed src=\"javascript:alert('XSS')\">", ""},
		{"<a href=\"http://example.com\" target=\"_blank\" rel=\"noopener noreferrer\">Safe Link</a>", "<a href=\"http://example.com\" rel=\"nofollow\">Safe Link</a>"},
		{"<a href=\"javascript:alert('XSS')\">Unsafe Link</a>", "Unsafe Link"},
		{"<style>@import 'javascript:alert(\"XSS\")';</style>", ""},
		{"<style>body { background-image: url('javascript:alert(\"XSS\")'); }</style>", ""},
		{"<script src=\"http://example.com/xss.js\"></script>", ""},
	}

	for i, s := range scenarios {
		t.Run(fmt.Sprintf("%d_%s", i, s.content), func(t *testing.T) {
			result := security.Sanitize(s.content)

			if s.expectContent != result {
				t.Fatalf("The sanitized string should be %s, got %s", s.expectContent, result)
			}

		})
	}
}
