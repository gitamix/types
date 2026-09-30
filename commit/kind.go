package commit

// Kind represents the kind of a commit.
//
// It is used to split commits by their kind in the linting process.
type Kind uint8

const (
	// KindDefault represents default commit.
	KindDefault Kind = 0
	// KindMerge represents merge commit.
	KindMerge Kind = 1
	// KindRevert represents revert commit.
	KindRevert Kind = 2
)

// Default defines whether the Kind is Default.
func (k Kind) Default() bool {
	return k == KindDefault
}

// Merge defines whether the Kind is Merge.
func (k Kind) Merge() bool {
	return k == KindMerge
}

// Revert defines whether the Kind is Revert.
func (k Kind) Revert() bool {
	return k == KindRevert
}
