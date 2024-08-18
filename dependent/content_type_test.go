package dependent

import "testing"

func TestParseContentType(t *testing.T) {
	ParseContentType([]byte(`text/plain; q=0.5`))
}
