package commit

import (
	"strings"
)

// Type represents the type of a commit subject.
//
// It is used to categorize the commit, such as "feat", "fix", etc.
type Type string

// NewType creates a new Type instance.
func NewType(value string) Type {
	return Type(value)
}

// Empty checks if the type is empty.
func (t Type) Empty() bool {
	return strings.TrimSpace(string(t)) == ""
}

// String returns the string representation of the Type.
func (t Type) String() string {
	return string(t)
}

// ParseType parses the type from a commit subject string.
//
// The expected format has the type at the beginning of the subject,
// optionally followed by a scope in parentheses and a colon.
//
// For example:
//
//	"feat(ui): add new button" -> type is "feat"
//	"fix: resolve issue"       -> type is "fix"
func ParseType(s string) Type {
	if len(s) == 0 {
		return NewType("")
	}
	ss := strings.SplitN(s, "(", 2)
	if len(ss) == 2 {
		return NewType(strings.TrimSpace(ss[0]))
	}
	ss = strings.SplitN(s, ":", 2)
	if len(ss) == 2 {
		return NewType(strings.TrimSpace(ss[0]))
	}
	return NewType("")
}
