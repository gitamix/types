package commit

// Commit represents a git commit in the version control system.
type Commit struct {
	// msg is the commit message associated with the commit.
	msg Message
	// hash is the commit hash defined after the commit is created.
	hash Hash

	kind Kind
}

// NewCommit creates a new Commit instance.
//
//   - hash: the commit hash.
//   - msg: the commit message.
//   - opts: functional options to configure the instance.
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

// Kind returns the kind of the commit.
//
// If no kind is provided in the constructor,
// it returns the default KindDefault.
func (c Commit) Kind() Kind {
	return c.kind
}
