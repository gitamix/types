package commit

import (
	"bytes"
	"regexp"
	"strings"

	"github.com/gitamix/types/ticket"
)

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
func ParseSubject[T string | []byte](v T) Subject {
	var s string
	switch val := any(v).(type) {
	case string:
		if i := strings.IndexRune(val, '\n'); i != -1 {
			s = val[:i]
		} else {
			s = val
		}
	case []byte:
		if i := bytes.IndexByte(val, '\n'); i != -1 {
			s = string(val[:i])
		} else {
			s = string(val)
		}
	}
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

// Ticket extracts a Ticket from the Subject using the provided regular expression.
//
// The regular expression should contain a capturing group that matches the ticket name.
// If the subject is empty or does not match the regular expression, an empty Ticket is returned.
//
// Example usage:
//
//	re := regexp.MustCompile(`^((?:TASK|PROJ|BUG)-\d+)`)
//	t := ParseSubject("TASK-1234 add new feature", re).Ticket(re)
//	fmt.Println(t.Name()) // Output: TASK-1234
func (s Subject) Ticket(re *regexp.Regexp) ticket.Ticket {
	return ticket.ParseTicket(s.String(), re)
}
