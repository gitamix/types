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
	return strings.TrimSpace(t.String()) == ""
}

// String returns the string representation of the Type.
func (t Type) String() string {
	return string(t)
}

// ParseType parses the type from a commit subject string.
//
// The expected format has the type at the beginning of the subject,
// optionally followed by a scope in parentheses and a colon.
// A leading ticket prefix (such as "WS-1234") is stripped automatically.
//
// For example:
//
//	"feat(ui): add new button"        -> type is "feat"
//	"[WS-1234] feat(ui): some feature"  -> type is "feat"
//	"fix: resolve issue"              -> type is "fix"
//	"feat"                           -> type is ""
func ParseType(s string) Type {
	if len(s) == 0 {
		return ""
	}
	i := strings.IndexAny(s, "(:")
	if i < 0 {
		return ""
	}
	v := strings.TrimSpace(s[:i])
	if fields := strings.Fields(v); len(fields) > 1 {
		v = fields[len(fields)-1]
	}
	return NewType(strings.TrimSpace(v))
}
