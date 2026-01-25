package branch

// Branch represents a git branch with its name.
//
// It encapsulates the branch name and provides methods to interact with it.
type Branch struct {
	// name is the name of the git branch.
	name Name
}

// NewBranch creates a new Branch instance with the given Name.
func NewBranch(name Name) Branch {
	return Branch{
		name: name,
	}
}

// String returns the string representation of the Branch, which is its Name.
func (b Branch) String() string {
	return b.name.String()
}
