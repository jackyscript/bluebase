package jsvm

import (
	"bluebase/security"
	"testing"

	"github.com/dop251/goja"
)

func TestSecuritySanitizeBind(t *testing.T) {
	vm := goja.New()
	vm.Set("sanitize", security.Sanitize)

	_, err := vm.RunString(`
		const content = "<div onclick=\"alert('XSS')\">Content</div>"
		const sanitized = sanitize(content)

		if (sanitized != "<div>Content</div>") {
			throw new Error("Expected sanitized content be <div>Content</div>, got instead " + sanitized)
		}
	`)
	if err != nil {
		t.Fatal(err)
	}
}
