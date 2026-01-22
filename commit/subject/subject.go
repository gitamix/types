package subject

import "github.com/gitamix/types/commit"

// Subject represents the subject of a commit message.
type Subject struct {
	// typ is the type of the commit.
	typ commit.Type
	// scope is the scope of the commit.
	scope commit.Scope
	// description is the description of the commit.
	description Description
}

// NewSubject creates a new Subject instance.
//
//   - t: type of the commit.
//   - s: scope of the commit.
//   - d: description of the commit.
func NewSubject(
	t commit.Type,
	s commit.Scope,
	d Description,
) Subject {
	return Subject{
		typ:         t,
		scope:       s,
		description: d,
	}
}

// ParseSubject parses a commit subject string into a Subject instance.
//
// The expected format of the input string are:
//
//	"type: description"
//	"type(scope): description"
//	"description"
//
// Examples:
//
//	"feat: add new feature"
//	"fix(ui): resolve button issue"
//	"update documentation"
func ParseSubject(s string) Subject {
	return NewSubject(
		commit.ParseType(s),
		commit.ParseScope(s),
		ParseDescription(s),
	)
}
