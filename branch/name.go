package branch

// Name represents the name of a git branch.
type Name string

// NewName creates a new branch Name from the given string.
func NewName(name string) Name {
	return Name(name)
}

// Empty checks if the branch name is empty.
func (n Name) Empty() bool {
	return n == ""
}

// String returns the string representation of the branch name.
func (n Name) String() string {
	return string(n)
}
