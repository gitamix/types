package ticket

import "regexp"

// ParseTicket extracts a Ticket from the given string using the provided regular expression.
//
// The regular expression should contain a capturing group that matches the ticket name.
// If the string is empty or does not match the regular expression, an empty Ticket is returned.
//
// Example usage:
//
//	re := regexp.MustCompile(`JIRA-(\d+)`)
//	t := ParseTicket("Fix issue JIRA-1234", re)
//	fmt.Println(t.Name()) // Output: JIRA-1234
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
