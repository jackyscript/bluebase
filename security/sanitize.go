package security

import (
	"github.com/microcosm-cc/bluemonday"
)

// Sanitize sanitizes "content".
//
// It uses the bluemonday library to clean the HTML content using UGCPolicy.
// This policy allows a limited set of HTML tags and attributes that are generally safe for user-generated content.
// It is important to note that this function does not guarantee complete security against XSS attacks.
// Always validate and sanitize user input on the server side before processing it further.
func Sanitize(content string) string {
	p := bluemonday.UGCPolicy()
	return p.Sanitize(content)
}
