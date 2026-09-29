package commit

// Option configures a Commit instance during creation in the constructor.
type Option func(*Commit)

// WithKind returns an Option that sets the kind of the commit on its creation.
func WithKind(k Kind) Option {
	return func(c *Commit) {
		c.kind = k
	}
}
