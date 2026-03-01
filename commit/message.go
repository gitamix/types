package commit

import (
	"bytes"
	"regexp"

	"github.com/gitamix/types/ticket"
)

// Message represents a git commit message.
type Message struct {
	// subject is the subject line of the commit message.
	subject Subject
	// body is the body of the commit message.
	body Body
}

// NewMessage creates a new Message instance.
//
//   - subject: the subject of the commit message.
//   - body: the body of the commit message.
func NewMessage(
	subject Subject,
	body Body,
) Message {
	return Message{
		subject: subject,
		body:    body,
	}
}

// ParseMessage parses a commit message from a byte slice or string.
//
// It splits the input into subject and body at the first newline character.
// If there is a second newline immediately following the first (i.e., "subject\n\nbody"),
// that blank line is skipped so the body starts after it. If there is no newline, the entire
// input is treated as the subject, and the body is empty.
func ParseMessage[T []byte | string](v T) Message {
	var bb, subjbb, bodybb []byte
	var subji, bodyi int
	switch val := any(v).(type) {
	case string:
		bb = []byte(val)
	case []byte:
		bb = val
	}
	subji = bytes.IndexByte(bb, '\n')
	if subji == -1 {
		subjbb = bb
	} else {
		subjbb = bb[:subji]
		bodyi = subji + 1
	}
	if bodyi > 0 {
		lni := bytes.IndexByte(bb[bodyi:], '\n')
		if lni != -1 {
			bodyi += lni + 1
		}
		bodybb = bb[bodyi:]
	}
	return NewMessage(
		ParseSubject(subjbb),
		NewBody(bodybb),
	)
}

// String returns the string representation of the Message.
//
// If the body is empty, it returns only the subject.
// Otherwise, it concatenates the subject and body with two newlines in between.
func (m Message) String() string {
	if m.body.Empty() {
		return m.subject.String()
	}
	return m.subject.String() + "\n\n" + m.body.String()
}

// Subject returns the subject of the commit message.
func (m Message) Subject() Subject {
	return m.subject
}

// Body returns the body of the commit message.
func (m Message) Body() Body {
	return m.body
}

// Ticket extracts a Ticket from the Message's subject using the provided regular expression.
//
// The regular expression should contain a capturing group that matches the ticket name.
// If the subject is empty or does not match the regular expression, an empty Ticket is returned.
//
// If there is no ticket found in the subject, it will not attempt to search the body for a ticket.
// Only the subject is considered for ticket extraction.
//
// Example usage:
//
//	re := regexp.MustCompile(`^((?:TASK|PROJ|BUG)-\d+)`)
//	t := ParseMessage("TASK-1234 add new feature", re).Ticket(re)
//	fmt.Println(t.Name()) // Output: TASK-1234
func (m Message) Ticket(re *regexp.Regexp) ticket.Ticket {
	return m.subject.Ticket(re)
}
