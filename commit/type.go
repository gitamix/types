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
