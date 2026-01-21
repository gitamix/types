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
