package commit

// Hash represents a commit hash in the version control system.
//
// It uniquely identifies a commit, regardless of the underlying hash algorithm.
type Hash string

// NewHash creates a new Hash instance.
func NewHash(value string) Hash {
	return Hash(value)
}

// String returns the string representation of the commit hash.
func (h Hash) String() string {
	return string(h)
}

// ShortString returns a shortened version of the commit hash.
func (h Hash) ShortString() string {
	if len(h) < 7 {
		return string(h)
	}
	return string(h)[:7]
}

// Empty checks if the commit hash is empty.
func (h Hash) Empty() bool {
	return len(h) == 0
}
