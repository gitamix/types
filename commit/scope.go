package commit

import (
	"strings"
)

// Scope represents the scope of a commit.
//
// It is typically used to specify the area of the codebase that the commit affects.
// For example, "(core)", "(ui)", etc.
type Scope string

// NewScope creates a new Scope instance.
func NewScope(value string) Scope {
	return Scope(value)
}

// Empty checks if the scope is empty.
func (s Scope) Empty() bool {
	return strings.TrimSpace(string(s)) == ""
}

// String returns the string representation of the Scope.
func (s Scope) String() string {
	return string(s)
}

// ParseScope parses the scope from a commit subject string.
//
// The expected format has the scope enclosed in parentheses.
//
// For example:
//
//	"feat(ui): add new button" -> scope is "ui"
func ParseScope(s string) Scope {
	if len(s) == 0 {
		return NewScope("")
	}
	ss := strings.SplitN(s, "(", 2)
	if len(ss) < 2 {
		return NewScope("")
	}
	if !strings.Contains(ss[1], ")") {
		return NewScope("")
	}
	ss = strings.SplitN(ss[1], ")", 2)
	return NewScope(
		strings.TrimSpace(ss[0]),
	)
}
