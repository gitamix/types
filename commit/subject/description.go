package subject

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
