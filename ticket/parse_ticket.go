package ticket

import "regexp"

// ParseTicket extracts a Ticket from the given string using the provided regular expression.
//
// The regular expression should contain a capturing group that matches the ticket name.
// If the string is empty or does not match the regular expression, an empty Ticket is returned.
//
// Example usage:
//
//	re := regexp.MustCompile(`^((?:TASK|PROJ|BUG)-\d+)`)
//	t := ParseTicket("TASK-1234 add new feature", re)
//	fmt.Println(t.Name()) // Output: TASK-1234
func ParseTicket(s string, re *regexp.Regexp) Ticket {
	if len(s) == 0 {
		return Ticket{}
	}
	matches := re.FindStringSubmatch(s)
	if len(matches) < 2 {
		return Ticket{}
	}
	return NewTicket(NewName(matches[1]))
}
