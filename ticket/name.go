package ticket

// Name represents the name of a task tracker ticket.
// It is a string that identifies the ticket and is used for display purposes.
type Name string

// NewName creates a new Name instance from a given string.
func NewName(name string) Name {
	return Name(name)
}

// String returns the string representation of the Name.
func (n Name) String() string {
	return string(n)
}

// Equals checks if two Name instances are equal.
func (n Name) Equals(other Name) bool {
	return n == other
}

// Empty checks if the Name is empty.
func (n Name) Empty() bool {
	return n == ""
}
