package commit

// Commit represents a git commit in the version control system.
type Commit struct {
	// msg is the commit message associated with the commit.
	msg Message
	// hash is the commit hash defined after the commit is created.
	hash Hash
	// merged tells if the commit is a merge commit with more than one parent.
	merged bool
}

// NewCommit creates a new Commit instance.
//
//   - hash: the commit hash.
//   - msg: the commit message.
//   - opts: the functional options to configure the Commit.
func NewCommit(
	hash Hash,
	msg Message,
	opts ...Option,
) Commit {
	c := Commit{
		hash: hash,
		msg:  msg,
	}
	for _, opt := range opts {
		opt(&c)
	}
	return c
}

// String returns the short string representation of the commit hash.
func (c Commit) String() string {
	return c.hash.ShortString()
}

// Message returns the commit message associated with the commit.
func (c Commit) Message() Message {
	return c.msg
}

// Hash returns the commit hash.
func (c Commit) Hash() Hash {
	return c.hash
}

// Merged reports whether the commit is a merge commit,
// i.e. a commit with more than one parent.
func (c Commit) Merged() bool {
	return c.merged
}

// Type returns the type of the commit
// as defined in the subject of the commit message.
//
// If no type is defined, it returns an empty Type.
//
// Example:
//
//	"feat(ui): add new button" -> "feat"
//	"fix: resolve issue"       -> "fix"
func (c Commit) Type() Type {
	return c.msg.subject.typ
}

// Scope returns the scope of the commit
// as defined in the subject of the commit message.
//
// If no scope is defined, it returns an empty Scope.
//
// Example:
//
//	"feat(ui): add new button" -> "ui"
//	"fix: resolve issue"       -> ""
func (c Commit) Scope() Scope {
	return c.msg.subject.scope
}
