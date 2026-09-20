package commit

// Option defines a functional option to configure
// the Commit instance created by the constructor.
type Option func(*Commit)

// WithMerged rmarks the commit as a merge commit on its creation.
//
// A merge commit is a commit with more than one parent,
// created by merging one branch into another.
func WithMerged() Option {
	return func(c *Commit) {
		c.merged = true
	}
}
