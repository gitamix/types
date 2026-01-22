package commit

import (
	"bytes"
)

// Body represents the body of a commit message.
//
// It is a byte slice that may contain multiple lines of text.
type Body []byte

// NewBody creates a new Body instance from the given byte slice.
//
//   - bb: byte slice representing the commit body.
func NewBody(bb []byte) Body {
	return Body(bb)
}

// Empty checks if the Body is empty or contains only whitespace.
func (b Body) Empty() bool {
	if len(b) == 0 {
		return true
	}
	return len(bytes.TrimSpace(b)) == 0
}

// String returns the string representation of the Body.
func (b Body) String() string {
	return string(b)
}
