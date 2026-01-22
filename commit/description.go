package commit

import "strings"

// Description represents the description part of a commit subject.
type Description string

// NewDescription creates a new Description instance.
func NewDescription(value string) Description {
	return Description(value)
}

// Empty checks if the description is empty.
func (d Description) Empty() bool {
	return strings.TrimSpace(string(d)) == ""
}

// String returns the string representation.
func (d Description) String() string {
	return string(d)
}

// ParseDescription parses the description from a commit subject string.
//
// The expected format has the description following a colon.
//
// For example:
//
//	"feat(ui): add new button" -> description is "add new button"
func ParseDescription(s string) Description {
	if len(s) == 0 {
		return NewDescription(s)
	}
	ss := strings.SplitN(s, ":", 2)
	if len(ss) < 2 {
		return NewDescription(s)
	}
	return NewDescription(
		strings.TrimSpace(ss[len(ss)-1]),
	)
}
