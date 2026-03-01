package branch

import (
	"regexp"
	"strings"

	"github.com/gitamix/types/ticket"
)

// Name represents the name of a git branch.
type Name string

// NewName creates a new branch Name from the given string.
func NewName(name string) Name {
	return Name(name)
}

// Empty checks if the branch name is empty.
func (n Name) Empty() bool {
	return strings.TrimSpace(string(n)) == ""
}

// String returns the string representation of the branch name.
func (n Name) String() string {
	return string(n)
}

// Ticket extracts a Ticket from the branch name using the provided regular expression.
//
// The regular expression should contain a capturing group that matches the ticket name.
// If the branch name is empty or does not match the regular expression, an empty Ticket is returned.
//
// Example usage:
//
//	re := regexp.MustCompile(`^(?:feature|bugfix|hotfix)/([A-Z]+-\d+)`)
//	t := NewName("feature/TASK-1234").Ticket(re)
//	fmt.Println(t.Name()) // Output: TASK-1234
func (n Name) Ticket(re *regexp.Regexp) ticket.Ticket {
	return ticket.ParseTicket(n.String(), re)
}
