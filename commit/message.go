package commit

import (
	"bytes"
	"regexp"

	"github.com/gitamix/types/ticket"
)

// MessageOption configures a Message at construction time.
//
// It is passed as an optional argument to NewMessage and ParseMessage
// to set fields that are not part of the required subject and body,
// such as the raw bytes the message was parsed from.
type MessageOption func(*Message)

// WithRaw returns a MessageOption that sets the raw bytes of a Message.
//
// The raw value preserves the original bytes the message was parsed from
// so that methods such as Ticket can inspect the unmodified input.
func WithRaw(raw []byte) MessageOption {
	return func(m *Message) {
		m.raw = raw
	}
}

// Message represents a git commit message.
//
// A Message is composed of a Subject and a Body, and optionally stores
// the raw bytes it was parsed from. Use NewMessage to build one from its
// parts, or ParseMessage to build one from a raw string or byte slice.
type Message struct {
	// raw stores the original bytes the message was parsed from.
	//
	// The value is preserved so that methods such as Ticket can inspect
	// the unmodified input. It is nil for messages built with NewMessage
	// unless WithRaw is supplied.
	raw []byte

	// subject is the subject line of the commit message.
	subject Subject

	// body is the body of the commit message.
	body Body
}

// NewMessage creates a new Message instance.
//
//   - subject: the subject of the commit message.
//   - body: the body of the commit message.
//   - opts: optional configuration applied to the Message, such as WithRaw.
func NewMessage(
	subject Subject,
	body Body,
	opts ...MessageOption,
) Message {
	m := Message{
		subject: subject,
		body:    body,
	}
	for _, opt := range opts {
		opt(&m)
	}
	return m
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
	m := NewMessage(
		ParseSubject(subjbb),
		NewBody(bodybb),
		WithRaw(bb),
	)
	return m
}

// String returns the string representation of the Message.
//
// When the Message stores raw bytes (for example one produced by ParseMessage,
// or by NewMessage with the WithRaw option), the raw bytes are returned as-is,
// preserving the original formatting of the commit message.
//
// Otherwise the representation is rebuilt from the subject and body: if the body
// is empty, only the subject is returned; otherwise the subject and body are
// joined with two newlines in between.
func (m Message) String() string {
	if len(m.raw) > 0 {
		return string(m.raw)
	}
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

// Ticket extracts a Ticket from the subject of the Message using the
// provided regular expression.
//
// The regular expression should contain a capturing group that matches
// the ticket name. Only the subject is considered for ticket extraction;
// the body is never searched.
//
// The subject is read from the raw bytes stored on the Message (the part
// before the first newline) rather than from the parsed Subject. This
// preserves a leading ticket prefix that ParseSubject strips, so a subject
// such as "TASK-1234 fix(ui): add button" still yields "TASK-1234" even
// though the parsed Subject no longer contains it.
//
// Ticket returns an empty Ticket when the Message has no raw bytes or when
// the raw bytes contain no newline character. In particular, a Message
// built with NewMessage without the WithRaw option always yields an empty
// Ticket. Use ParseMessage, which records the raw bytes, to ensure
// extraction works as expected.
//
// Panics if re is nil.
//
// Example usage:
//
//	re := regexp.MustCompile(`^((?:TASK|PROJ|BUG)-\d+)`)
//	t := ParseMessage("TASK-1234 add new feature\n\nDetails here.").Ticket(re)
//	fmt.Println(t.Name()) // Output: TASK-1234
func (m Message) Ticket(re *regexp.Regexp) ticket.Ticket {
	subji := bytes.IndexByte(m.raw, '\n')
	if subji == -1 {
		return ticket.Ticket{}
	}
	subjbb := m.raw[:subji]
	return ticket.ParseTicket(string(subjbb), re)
}
