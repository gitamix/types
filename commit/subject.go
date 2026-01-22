package commit

// Subject represents the subject of a commit message.
type Subject struct {
	// typ is the type of the commit.
	typ Type
	// scope is the scope of the commit.
	scope Scope
	// description is the description of the commit.
	description Description
}

// NewSubject creates a new Subject instance.
//
//   - t: type of the commit.
//   - s: scope of the commit.
//   - d: description of the commit.
func NewSubject(
	t Type,
	s Scope,
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
// The expected formats of the input string are:
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
		ParseType(s),
		ParseScope(s),
		ParseDescription(s),
	)
}

// String returns the string representation of the Subject.
//
// The format is:
//
//	"type(scope): description"
//	"type: description"
//	"description"
//
// depending on the presence of type and scope.
//
// Examples:
//
//	"feat(ui): add new feature"
//	"fix: resolve button issue"
//	"update documentation"
//
// Note: If both type and scope are empty, only the description is returned.
func (s Subject) String() string {
	if s.scope.Empty() {
		if s.typ.Empty() {
			return s.description.String()
		}
		return s.typ.String() + ": " + s.description.String()
	}
	if s.typ.Empty() {
		return s.description.String()
	}
	return s.typ.String() + "(" + s.scope.String() + "): " + s.description.String()
}
